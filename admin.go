package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"

	"time"

	"github.com/gorilla/mux"
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
		return usersinfo[i].Lastname < usersinfo[j].Lastname
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
