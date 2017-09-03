package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gopkg.in/authboss.v1"
	_ "gopkg.in/authboss.v1/auth"
	_ "gopkg.in/authboss.v1/confirm"
	_ "gopkg.in/authboss.v1/lock"
	aboauth "gopkg.in/authboss.v1/oauth2"
	_ "gopkg.in/authboss.v1/recover"
	_ "gopkg.in/authboss.v1/register"
	_ "gopkg.in/authboss.v1/remember"

	"encoding/json"

	"github.com/aarondl/tpl"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/justinas/alice"
	"github.com/justinas/nosurf"
	sprockets "github.com/znly/go-sprockets"
)

var funcs = template.FuncMap{
	"formatDate": func(date time.Time) string {
		return date.Format("2006/01/02 03:04pm")
	},
	"yield": func() string { return "" },
}

var (
	ab        = authboss.New()
	sprocket  *sprockets.Sprocket
	templates = tpl.Must(tpl.Load("views", "views/partials", "layout.html.tpl", funcs))
	schemaDec = schema.NewDecoder()
	db        *gorm.DB
	serverURL = os.Getenv("SERVER_URL")
)

func setupAuthboss(database *DatabaseStorer) {
	ab.Storer = database
	ab.OAuth2Storer = database
	ab.MountPath = "/auth"
	ab.ViewsPath = "ab_views"
	ab.RootURL = serverURL

	ab.LayoutDataMaker = layoutData

	ab.OAuth2Providers = map[string]authboss.OAuth2Provider{
		"google": authboss.OAuth2Provider{
			OAuth2Config: &oauth2.Config{
				ClientID:     ``,
				ClientSecret: ``,
				Scopes:       []string{`profile`, `email`},
				Endpoint:     google.Endpoint,
			},
			Callback: aboauth.Google,
		},
	}

	b, err := ioutil.ReadFile(filepath.Join("views", "layout.html.tpl"))
	if err != nil {
		panic(err)
	}
	ab.Layout = template.Must(template.New("layout").Funcs(funcs).Parse(string(b)))

	ab.XSRFName = "csrf_token"
	ab.XSRFMaker = func(_ http.ResponseWriter, r *http.Request) string {
		return nosurf.Token(r)
	}

	ab.CookieStoreMaker = NewCookieStorer
	ab.SessionStoreMaker = NewSessionStorer
	ab.EmailFrom = "Fall Of Summer Benevole <noreply-benevole-fos@korporation.net>"
	ab.Mailer = authboss.SMTPMailer("ssl0.ovh.net:587", smtp.PlainAuth("noreply-benevole-fos@korporation.net", "noreply-benevole-fos@korporation.net", os.Getenv("MAIL_PASSWORD"), "ssl0.ovh.net"))
	//ab.Mailer = authboss.LogMailer(os.Stdout)

	ab.Policies = []authboss.Validator{
		authboss.Rules{
			FieldName:       "email",
			Required:        true,
			AllowWhitespace: false,
		},
		authboss.Rules{
			FieldName:       "password",
			Required:        true,
			MinLength:       4,
			MaxLength:       30,
			AllowWhitespace: false,
		},
	}

	if err := ab.Init(); err != nil {
		log.Fatal(err)
	}
}

func setupAssetManager() {
	var err error
	sprocket, err = sprockets.NewWithDefault("./assets", "")
	if err != nil {
		log.Fatal(err)
	}
	err = sprocket.PushFrontExtensionPath(".jpg", os.Getenv("PHOTO_PATH"))
	if err != nil {
		log.Fatal(err)
	}
}

func ConvertFormDate(value string) reflect.Value {
	s, _ := time.Parse("2006-01-_2", value)
	return reflect.ValueOf(s)
}

func main() {
	// Initialize Sessions and Cookies
	// Typically gorilla securecookie and sessions packages require
	// highly random secret keys that are not divulged to the public.
	//
	// In this example we use keys generated one time (if these keys ever become
	// compromised the gorilla libraries allow for key rotation, see gorilla docs)
	// The keys are 64-bytes as recommended for HMAC keys as per the gorilla docs.
	//
	// These values MUST be changed for any new project as these keys are already "compromised"
	// as they're in the public domain, if you do not change these your application will have a fairly
	// wide-opened security hole. You can generate your own with the code below, or using whatever method
	// you prefer:
	//
	//    func main() {
	//        fmt.Println(base64.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(64)))
	//    }
	//
	// We store them in base64 in the example to make it easy if we wanted to move them later to
	// a configuration environment var or file.
	cookieStoreKey, _ := base64.StdEncoding.DecodeString(os.Getenv("COOKIE_STORE_KEY"))
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(os.Getenv("SESSION_STORE_KEY"))
	cookieStore = securecookie.New(cookieStoreKey, nil)
	sessionStore = sessions.NewCookieStore(sessionStoreKey)

	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE")))
	if err != nil {
		log.Fatalln(err)
	}
	database := NewDataBaseStorer(db)
	// Initialize ab.
	setupAuthboss(database)

	// Setup the asset manager
	setupAssetManager()

	// Set up our router
	schemaDec.IgnoreUnknownKeys(true)
	schemaDec.RegisterConverter(time.Time{}, ConvertFormDate)
	mux := mux.NewRouter()

	// Routes
	gets := mux.Methods("GET").Subrouter()

	mux.PathPrefix("/auth").Handler(ab.NewRouter())
	gets.PathPrefix("/assets/").HandlerFunc(assetsHandler)
	gets.HandleFunc("/", index)
	gets.Handle("/form", authProtect(form))

	gets.Handle("/admin", authProtect(admin))
	gets.Handle("/admin/csv", authProtect(generateCSV))
	gets.Handle("/admin/pdf/{id}", authProtect(generateOnePDF))
	gets.Handle("/admin/allpdf", authProtect(generateAllPDF))

	gets.Handle("/setadmin123/{id}", authProtect(setadmin))

	// posts := mux.Methods("POST").Subrouter()
	//posts.Handle("/upload", authProtect(uploadPhoto))
	//posts.Handle("/form", authProtect(formPost))

	mux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not found")
	})

	// Set up our middleware chain
	stack := alice.New(logger, nosurfing, ab.ExpireMiddleware).Then(mux)

	// Start the server
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	log.Println(http.ListenAndServe(":"+port, stack))
}

func layoutData(w http.ResponseWriter, r *http.Request) authboss.HTMLData {
	currentUserName := ""
	birthdate := ""
	WhatYouWantToDo := ""
	picture := ""
	var userInfo UserInfo
	var user *User
	userInter, err := ab.CurrentUser(w, r)
	if userInter != nil && err == nil {
		user = userInter.(*User)
		currentUserName = user.Name
		db.First(&userInfo, "id = ?", user.ID)
		birthdate = userInfo.BirthDate.Format("2006-01-02")
		TmpWhatYouWantToDo := make([]JobsType, 0, 4)
		if userInfo.WhatYouWantToDo1 > 0 {
			TmpWhatYouWantToDo = append(TmpWhatYouWantToDo, userInfo.WhatYouWantToDo1)
		}
		if userInfo.WhatYouWantToDo2 > 0 {
			TmpWhatYouWantToDo = append(TmpWhatYouWantToDo, userInfo.WhatYouWantToDo2)
		}
		if userInfo.WhatYouWantToDo3 > 0 {
			TmpWhatYouWantToDo = append(TmpWhatYouWantToDo, userInfo.WhatYouWantToDo3)
		}
		if userInfo.WhatYouWantToDo4 > 0 {
			TmpWhatYouWantToDo = append(TmpWhatYouWantToDo, userInfo.WhatYouWantToDo4)
		}
		if len(TmpWhatYouWantToDo) > 0 {
			if tmpVal, err := json.Marshal(&TmpWhatYouWantToDo); err == nil {
				WhatYouWantToDo = string(tmpVal)
			}
		}
		if len(userInfo.Photo) > 0 {
			picture = fmt.Sprintf("assets/%s.jpg", hex.EncodeToString([]byte(user.ID)))
		}
	}

	ret := authboss.HTMLData{
		"loggedin":               userInter != nil,
		"username":               "",
		authboss.FlashSuccessKey: ab.FlashSuccess(w, r),
		authboss.FlashErrorKey:   ab.FlashError(w, r),
		"current_user_name":      currentUserName,
		"birthdate":              birthdate,
		"what_you_want_to_do":    WhatYouWantToDo,
		"picture":                picture,
	}
	ret.MergeKV("userinfo", userInfo)
	return ret
}

func index(w http.ResponseWriter, r *http.Request) {
	data := layoutData(w, r)
	mustRender(w, r, "index", data)
}

func form(w http.ResponseWriter, r *http.Request) {
	data := layoutData(w, r)
	mustRender(w, r, "form", data)
}

func formPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if badRequest(w, err) {
		return
	}
	userInter, err := ab.CurrentUser(w, r)
	if internalError(w, err) {
		return
	}
	if userInter == nil {
		badRequest(w, fmt.Errorf("Not logged"))
		return
	}
	// TODO: Validation

	var userinfo UserInfo
	var olduserinfo UserInfo
	if db.First(&olduserinfo, "id = ?", userInter.(*User).ID).Error == nil {
		userinfo.CreatedAt = olduserinfo.CreatedAt
		if userinfo.CreatedAt.Before(time.Date(2017, time.May, 01, 00, 00, 01, 00, time.UTC)) {
			userinfo.CreatedAt = time.Now()
		}
	}
	if badRequest(w, schemaDec.Decode(&userinfo, r.PostForm)) {
		return
	}
	birthdateStr := r.PostForm["BirthDate"]
	if len(birthdateStr) == 0 {
		badRequest(w, fmt.Errorf("no birthdate"))
		return
	}
	birthdateDate, err := time.Parse("2006-01-02", birthdateStr[0])
	if err != nil {
		birthdateDate, err = time.Parse("2006-1-2", birthdateStr[0])
		if err != nil {
			badRequest(w, fmt.Errorf("bad birthdate"))
			return
		}
	}
	userinfo.BirthDate = birthdateDate

	var TmpWhatYouWantToDo []JobsType
	if whatYouWantToDo := r.PostForm["WhatYouWantToDo"]; len(whatYouWantToDo) == 0 {
		badRequest(w, fmt.Errorf("No jobs chosen"))
		return
	} else if badRequest(w, json.Unmarshal([]byte(r.PostForm["WhatYouWantToDo"][0]), &TmpWhatYouWantToDo)) {
		return
	} else if len(TmpWhatYouWantToDo) != 4 {
		badRequest(w, fmt.Errorf("Wrong number of jobs chosen"))
		return
	}
	userinfo.WhatYouWantToDo1 = TmpWhatYouWantToDo[0]
	userinfo.WhatYouWantToDo2 = TmpWhatYouWantToDo[1]
	userinfo.WhatYouWantToDo3 = TmpWhatYouWantToDo[2]
	userinfo.WhatYouWantToDo4 = TmpWhatYouWantToDo[3]
	userinfo.ID = userInter.(*User).ID

	if internalError(w, db.Save(&userinfo).Error) {
		return
	}
	http.Redirect(w, r, "/form?ok", http.StatusFound)
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/assets/"):]
	b, err := sprocket.GetAsset(path)
	switch err {
	case nil:
	case sprockets.ErrNotFound:
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-type", mime.TypeByExtension(filepath.Ext(path)))
	w.Write(b)
}

func mustRender(w http.ResponseWriter, r *http.Request, name string, data authboss.HTMLData) {
	data.MergeKV("csrf_token", nosurf.Token(r))
	err := templates.Render(w, name, data)
	if err == nil {
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, "Error occurred rendering template:", err)
}

func badRequest(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, "Bad request:", err)

	return true
}

func internalError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, "Internal Server Error:", err)

	return true
}
