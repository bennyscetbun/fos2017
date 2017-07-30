package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"time"

	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
)

var bennyEpoch = time.Date(2017, 05, 01, 00, 00, 00, 00, time.UTC)

var csvtab = []string{
	"Nom",
	"Prenom",
	"Email",
	"Adresse",
	"CP",
	"Ville",
	"Numero Secu",
	"Date de naissance",
	"Lieu de naissance",
	"Telephone",
	"Facebook",
	"TShirt",
	"Regime",
	"Allergy",
	"Info Medicale",
	"Permis VL",
	"Permis PL",
	"Premier Secours",
	"Niveau anglais",
	"Autre Langage",
	"Deja benevole FOS",
	"Deja benevole ailleurs",
	"deja venu au FOS",
	"Choix 1",
	"Choix 2",
	"Choix 3",
	"Choix 4",
	"Autre info Choix",
	//"Autre info",
	//"Photo",
	"Dimanche 3",
	"Lundi 4",
	"Mardi 5",
	"Mercredi 6",
	"Jeudi 7",
	"Vendredi 8",
	"Samedi 9",
	"Dimanche 10",
	"Lundi 11",
	"Mardi 12",
	"Mercredi 13",

	"Cree le",
	"Mis a jour le",
}

const dateFormat = "2006-01-02 15:04:05"

func admin(w http.ResponseWriter, r *http.Request) {
	userInter, err := ab.CurrentUser(w, r)
	if userInter != nil && err == nil {
		user := userInter.(*User)
		if user.IsAdmin == false {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	data := make(map[string]interface{})
	var usersinfo []UserInfo
	if internalError(w, db.Find(&usersinfo).Error) {
		return
	}
	data["usersinfo"] = usersinfo

	mustRender(w, r, "admin", data)
}

func setadmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userInter, err := ab.CurrentUser(w, r)
	if userInter != nil && err == nil {
		user := userInter.(*User)
		if user.IsAdmin == false {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	var user User
	if internalError(w, db.First(&user, "id = ?", vars["id"]).Error) {
		return
	}

	user.IsAdmin = true
	if internalError(w, db.Save(&user).Error) {
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func generateCSV(w http.ResponseWriter, r *http.Request) {
	userInter, err := ab.CurrentUser(w, r)
	if userInter != nil && err == nil {
		user := userInter.(*User)
		if user.IsAdmin == false {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	var usersinfo []UserInfo
	if internalError(w, db.Find(&usersinfo).Error) {
		return
	}
	buf := bytes.Buffer{}
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("1")
	if internalError(w, err) {
		return
	}
	{
		row := sheet.AddRow()
		for _, col := range csvtab {
			row.AddCell().SetValue(col)
		}
	}
	sort.Slice(usersinfo, func(i, j int) bool {
		return strings.ToLower(usersinfo[i].Lastname) < strings.ToLower(usersinfo[j].Lastname)
	})
	for _, userinfo := range usersinfo {
		row := sheet.AddRow()
		row.AddCell().SetValue(userinfo.Lastname)
		row.AddCell().SetValue(userinfo.Firstname)
		row.AddCell().SetValue(userinfo.ID)
		row.AddCell().SetValue(userinfo.Address)
		row.AddCell().SetValue(userinfo.CP)
		row.AddCell().SetValue(userinfo.Town)
		row.AddCell().SetValue(userinfo.HealthNumber)
		row.AddCell().SetValue(userinfo.BirthDate)
		row.AddCell().SetValue(userinfo.BirthPlace)
		row.AddCell().SetValue(userinfo.PhoneNumber)
		if userinfo.Facebook != nil {
			row.AddCell().SetValue(*userinfo.Facebook)
		} else {
			row.AddCell().SetValue("")
		}
		row.AddCell().SetValue(userinfo.TShirt.String())
		row.AddCell().SetValue(userinfo.Regime.String())
		row.AddCell().SetValue(userinfo.Allergy)
		row.AddCell().SetValue(userinfo.MedicalInfo)
		row.AddCell().SetValue(userinfo.DriverLicenceVL.String())
		row.AddCell().SetValue(userinfo.DriverLicencePL.String())
		row.AddCell().SetValue(userinfo.FirstAidTraining.String())
		row.AddCell().SetValue(userinfo.EnglishLevel.String())
		row.AddCell().SetValue(userinfo.OtherLanguage)
		row.AddCell().SetValue(userinfo.AlreadyBeenBenevolFOS)
		row.AddCell().SetValue(userinfo.AlreadyBeenBenevol)
		row.AddCell().SetValue(userinfo.DidYouCameFOS.String())
		row.AddCell().SetValue(userinfo.WhatYouWantToDo1.String())
		row.AddCell().SetValue(userinfo.WhatYouWantToDo2.String())
		row.AddCell().SetValue(userinfo.WhatYouWantToDo3.String())
		row.AddCell().SetValue(userinfo.WhatYouWantToDo4.String())
		row.AddCell().SetValue(userinfo.OtherJobs)
		//row.AddCell().SetValue(userinfo.OtherInfo)
		//row.AddCell().SetValue(fmt.Sprintf("%s/assets/%s.jpg", serverURL, hex.EncodeToString([]byte(userinfo.ID))))
		if userinfo.WhenCanYouBeThere&DayThereSunday1 != 0 {
			row.AddCell().SetValue("here")
		} else {
			row.AddCell().SetValue("")
		}
		if userinfo.WhenCanYouBeThere&DayThereMonday1 != 0 {
			row.AddCell().SetValue("here")
		} else {
			row.AddCell().SetValue("")
		}
		if userinfo.WhenCanYouBeThere&DayThereTuesday1 != 0 {
			row.AddCell().SetValue("here")
		} else {
			row.AddCell().SetValue("")
		}
		if userinfo.WhenCanYouBeThere&DayThereWesnesday1 != 0 {
			row.AddCell().SetValue("here")
		} else {
			row.AddCell().SetValue("")
		}
		if userinfo.WhenCanYouBeThere&DayThereThursday1 != 0 {
			row.AddCell().SetValue("here")
		} else {
			row.AddCell().SetValue("")
		}
		if userinfo.WhenCanYouBeThere&DayThereFriday1 != 0 {
			row.AddCell().SetValue("here")
		} else {
			row.AddCell().SetValue("")
		}
		if userinfo.WhenCanYouBeThere&DayThereSaturday1 != 0 {
			row.AddCell().SetValue("here")
		} else {
			row.AddCell().SetValue("")
		}
		if userinfo.WhenCanYouBeThere&DayThereSunday2 != 0 {
			row.AddCell().SetValue("here")
		} else {
			row.AddCell().SetValue("")
		}
		if userinfo.WhenCanYouBeThere&DayThereMonday2 != 0 {
			row.AddCell().SetValue("here")
		} else {
			row.AddCell().SetValue("")
		}
		if userinfo.WhenCanYouBeThere&DayThereTuesday2 != 0 {
			row.AddCell().SetValue("here")
		} else {
			row.AddCell().SetValue("")
		}
		if userinfo.WhenCanYouBeThere&DayThereWesnesday2 != 0 {
			row.AddCell().SetValue("here")
		} else {
			row.AddCell().SetValue("")
		}
		if userinfo.CreatedAt.Before(bennyEpoch) {
			row.AddCell().SetValue(bennyEpoch)
		} else {
			row.AddCell().SetValue(userinfo.CreatedAt)
		}
		if userinfo.UpdatedAt.Before(bennyEpoch) {
			row.AddCell().SetValue(bennyEpoch)
		} else {
			row.AddCell().SetValue(userinfo.UpdatedAt)
		}
	}
	if internalError(w, file.Write(&buf)) {
		return
	}
	w.Header().Set("Content-Type", "text/xlsx")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.xlsx\"", time.Now().Format("benevole_database_2006-01-02")))
	w.Write(buf.Bytes())
	return
}

func generatePDF(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userInter, err := ab.CurrentUser(w, r)
	if userInter != nil && err == nil {
		user := userInter.(*User)
		if user.IsAdmin == false {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	var userInfo UserInfo
	if internalError(w, db.First(&userInfo, "id = ?", vars["id"]).Error) {
		return
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pageWidth, _ := pdf.GetPageSize()
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	centerTextFct := func(text string, y, wOffset float64) float64 {
		title := tr(text)
		titleWidth := pdf.GetStringWidth(title)
		pdf.Text((pageWidth-wOffset-titleWidth)*0.5+wOffset, y, title)
		_, fontHeight := pdf.GetFontSize()
		return fontHeight + y
	}
	writeUnderFct := func(title string, text interface{}, y float64) float64 {
		_, fontHeight := pdf.GetFontSize()
		if len(title) > 0 {
			pdf.Text(10, y, tr(fmt.Sprintf("%s: %s", title, text)))
		} else {
			pdf.Text(10, y, tr(fmt.Sprintf("%s", text)))
		}
		return fontHeight + y + 2
	}

	header := func() float64 {
		pdf.AddPage()
		_, photoPath := localPhotoPath(userInfo.ID)
		pdf.RegisterImageOptions(photoPath, gofpdf.ImageOptions{})
		imgInfo := pdf.GetImageInfo(photoPath)
		if imgInfo.Height() > imgInfo.Width() {
			pdf.Image(photoPath, 10, 6, 0, 30, false, "", 0, "")
		} else {
			pdf.Image(photoPath, 10, 6, 30, 0, false, "", 0, "")
		}

		pdf.SetFont("Arial", "B", 15)

		y := 18.0
		y = centerTextFct(fmt.Sprintf("%s %s", userInfo.Lastname, userInfo.Firstname), y, 40)
		pdf.SetFont("Arial", "", 13)
		y += 3
		y = centerTextFct(fmt.Sprintf("Tel: %s", userInfo.PhoneNumber), y, 40)
		return 40 + 4
	}

	presence := func(y float64) float64 {
		pdf.SetFont("Arial", "B", 13)
		y = writeUnderFct("Presence", "", y)
		pdf.SetFont("Arial", "", 12)
		_, fontHeight := pdf.GetFontSize()
		x := 10.0
		for i, date := range []string{
			"Dimanche 3",
			"Lundi 4",
			"Mardi 5",
			"Mercredi 6",
			"Jeudi 7",
			"Vendredi 8",
			"Samedi 9",
			"Dimanche 10",
			"Lundi 11",
			"Mardi 12",
			"Mercredi 13",
		} {
			bit := DayThere(1) << uint(i)

			if i == 5 || i == 7 {
				x = 10.0
				y += (fontHeight+2)*2 + 5
			}
			pdf.SetXY(x, y)
			pdf.CellFormat(30, fontHeight+2, date, "1", 0, "CM", false, 0, "")
			pdf.SetXY(x, y+fontHeight+2)
			txt := ""
			if userInfo.WhenCanYouBeThere&bit != 0 {
				txt = "X"
			}
			pdf.CellFormat(30, fontHeight+2, txt, "1", 0, "CM", false, 0, "")
			x += 30
		}
		return y + (fontHeight+2)*2 + 5
	}

	y := header()
	pdf.SetFont("Arial", "B", 13)
	y = writeUnderFct("Etat Civil", "", y)
	pdf.SetFont("Arial", "", 12)
	y = writeUnderFct("Date de naissance/lieu de naissance", fmt.Sprintf("%s/%s", userInfo.BirthDate.Format("02-01-2006"), userInfo.BirthPlace), y)
	y = writeUnderFct("Numero Secu", userInfo.HealthNumber, y)
	y = writeUnderFct("Adresse", fmt.Sprintf("%s %s %s", userInfo.Address, userInfo.CP, userInfo.Town), y)
	y = writeUnderFct("email", userInfo.ID, y)
	y = writeUnderFct("Tshirt", userInfo.TShirt, y)
	y += 5
	pdf.SetFont("Arial", "B", 13)
	y = writeUnderFct("Competence", "", y)
	pdf.SetFont("Arial", "", 12)
	y = writeUnderFct("Permis VL", userInfo.DriverLicenceVL, y)
	y = writeUnderFct("Permis PL", userInfo.DriverLicencePL, y)
	y = writeUnderFct("Premier Secours", userInfo.FirstAidTraining, y)
	y = writeUnderFct("Anglais", userInfo.EnglishLevel, y)

	y = writeUnderFct("Autres langues parlées", userInfo.OtherLanguage, y)

	y += 5
	y = presence(y)
	y += 5
	pdf.SetFont("Arial", "B", 13)
	y = writeUnderFct("Contact D urgence", "", y)
	pdf.SetFont("Arial", "", 12)
	y = writeUnderFct("Lien", userInfo.EmergencyContactType, y)
	y = writeUnderFct("Nom", userInfo.EmergencyContactLastname, y)
	y = writeUnderFct("Prenom", userInfo.EmergencyContactFirstname, y)
	y = writeUnderFct("Tel", userInfo.EmergencyContactPhoneNumber, y)
	y = writeUnderFct("Adresse", fmt.Sprintf("%s %s %s", userInfo.EmergencyContactAddress, userInfo.EmergencyContactCP, userInfo.EmergencyContactTown), y)
	y += 5
	pdf.SetFont("Arial", "B", 13)
	y = writeUnderFct("Info Medicale", "", y)
	pdf.SetFont("Arial", "", 12)
	y = writeUnderFct("Allergies", userInfo.Allergy, y)
	y = writeUnderFct("Regime alimentaire", userInfo.Regime, y)

	//html.
	/*
		//

		//
		pdf.SetFont("Arial", "B", 12)

		for _, str := range []string{"Nom", "Prenom", "Date de naissance"} {
			pdf.CellFormat(40, 7, str, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(40, 6, userInfo.Lastname, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 6, userInfo.Firstname, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 6, userInfo.BirthDate.Format("01-02-2006"), "1", 0, "", false, 0, "")
		pdf.Ln(-1)

	*/
	buf := bytes.Buffer{}
	if internalError(w, pdf.Output(&buf)) {
		return
	}
	w.Header().Set("Content-Type", "text/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s_%s.pdf\"", userInfo.Firstname, userInfo.Lastname))
	w.Write(buf.Bytes())
	//err := pdf.OutputFileAndClose("hello.pdf")
}
