package main

import (
	"bytes"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"net/http"

	"time"

	"github.com/gorilla/mux"
	"github.com/tealeg/xlsx"
)

var bennyEpoch = time.Date(2017, 05, 01, 00, 00, 00, 00, time.UTC)

var csvtab = []string{
	"Email",
	"Nom",
	"Prenom",
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
	"Autre info",
	"Photo",
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
	var b bytes.Buffer
	csvWriter := csv.NewWriter(&b)
	csvWriter.Comma = ';'
	if internalError(w, csvWriter.Write(csvtab)) {
		return
	}
	for _, userinfo := range usersinfo {
		toWrite := make([]string, 0, len(csvtab))
		toWrite = append(toWrite, userinfo.ID)
		toWrite = append(toWrite, userinfo.Lastname)
		toWrite = append(toWrite, userinfo.Firstname)
		toWrite = append(toWrite, userinfo.Address)
		toWrite = append(toWrite, userinfo.CP)
		toWrite = append(toWrite, userinfo.Town)
		toWrite = append(toWrite, userinfo.HealthNumber)
		toWrite = append(toWrite, userinfo.BirthDate.Format(dateFormat))
		toWrite = append(toWrite, userinfo.BirthPlace)
		toWrite = append(toWrite, userinfo.PhoneNumber)
		if userinfo.Facebook != nil {
			toWrite = append(toWrite, *userinfo.Facebook)
		} else {
			toWrite = append(toWrite, "")
		}
		toWrite = append(toWrite, userinfo.TShirt.String())
		toWrite = append(toWrite, userinfo.Regime.String())
		toWrite = append(toWrite, userinfo.Allergy)
		toWrite = append(toWrite, userinfo.MedicalInfo)
		toWrite = append(toWrite, userinfo.DriverLicenceVL.String())
		toWrite = append(toWrite, userinfo.DriverLicencePL.String())
		toWrite = append(toWrite, userinfo.FirstAidTraining.String())
		toWrite = append(toWrite, userinfo.EnglishLevel.String())
		toWrite = append(toWrite, userinfo.OtherLanguage)
		toWrite = append(toWrite, userinfo.AlreadyBeenBenevolFOS)
		toWrite = append(toWrite, userinfo.AlreadyBeenBenevol)
		toWrite = append(toWrite, userinfo.DidYouCameFOS.String())
		toWrite = append(toWrite, userinfo.WhatYouWantToDo1.String())
		toWrite = append(toWrite, userinfo.WhatYouWantToDo2.String())
		toWrite = append(toWrite, userinfo.WhatYouWantToDo3.String())
		toWrite = append(toWrite, userinfo.WhatYouWantToDo4.String())
		toWrite = append(toWrite, userinfo.OtherJobs)
		toWrite = append(toWrite, userinfo.OtherInfo)
		toWrite = append(toWrite, fmt.Sprintf("%s/assets/%s.jpg", serverURL, hex.EncodeToString([]byte(userinfo.ID))))
		if userinfo.WhenCanYouBeThere&DayThereSunday1 != 0 {
			toWrite = append(toWrite, "here")
		} else {
			toWrite = append(toWrite, "")
		}
		if userinfo.WhenCanYouBeThere&DayThereMonday1 != 0 {
			toWrite = append(toWrite, "here")
		} else {
			toWrite = append(toWrite, "")
		}
		if userinfo.WhenCanYouBeThere&DayThereTuesday1 != 0 {
			toWrite = append(toWrite, "here")
		} else {
			toWrite = append(toWrite, "")
		}
		if userinfo.WhenCanYouBeThere&DayThereWesnesday1 != 0 {
			toWrite = append(toWrite, "here")
		} else {
			toWrite = append(toWrite, "")
		}
		if userinfo.WhenCanYouBeThere&DayThereThursday1 != 0 {
			toWrite = append(toWrite, "here")
		} else {
			toWrite = append(toWrite, "")
		}
		if userinfo.WhenCanYouBeThere&DayThereFriday1 != 0 {
			toWrite = append(toWrite, "here")
		} else {
			toWrite = append(toWrite, "")
		}
		if userinfo.WhenCanYouBeThere&DayThereSaturday1 != 0 {
			toWrite = append(toWrite, "here")
		} else {
			toWrite = append(toWrite, "")
		}
		if userinfo.WhenCanYouBeThere&DayThereSunday2 != 0 {
			toWrite = append(toWrite, "here")
		} else {
			toWrite = append(toWrite, "")
		}
		if userinfo.WhenCanYouBeThere&DayThereMonday2 != 0 {
			toWrite = append(toWrite, "here")
		} else {
			toWrite = append(toWrite, "")
		}
		if userinfo.WhenCanYouBeThere&DayThereTuesday2 != 0 {
			toWrite = append(toWrite, "here")
		} else {
			toWrite = append(toWrite, "")
		}
		if userinfo.WhenCanYouBeThere&DayThereWesnesday2 != 0 {
			toWrite = append(toWrite, "here")
		} else {
			toWrite = append(toWrite, "")
		}
		if userinfo.CreatedAt.Before(bennyEpoch) {
			toWrite = append(toWrite, bennyEpoch.Format(dateFormat))
		} else {
			toWrite = append(toWrite, userinfo.CreatedAt.Format(dateFormat))
		}
		if userinfo.UpdatedAt.Before(bennyEpoch) {
			toWrite = append(toWrite, bennyEpoch.Format(dateFormat))
		} else {
			toWrite = append(toWrite, userinfo.UpdatedAt.Format(dateFormat))
		}
		if internalError(w, csvWriter.Write(toWrite)) {
			return
		}
	}
	csvWriter.Flush()
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.csv\"", time.Now().Format("benevole_database_2006-01-02")))
	w.Write(b.Bytes())
	return
}

func generateCSV2(w http.ResponseWriter, r *http.Request) {
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
	for _, userinfo := range usersinfo {
		row := sheet.AddRow()
		row.AddCell().SetValue(userinfo.ID)
		row.AddCell().SetValue(userinfo.Lastname)
		row.AddCell().SetValue(userinfo.Firstname)
		row.AddCell().SetValue(userinfo.Address)
		row.AddCell().SetValue(userinfo.CP)
		row.AddCell().SetValue(userinfo.Town)
		row.AddCell().SetValue(userinfo.HealthNumber)
		row.AddCell().SetValue(userinfo.BirthDate.Format(dateFormat))
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
		row.AddCell().SetValue(userinfo.OtherInfo)
		row.AddCell().SetValue(fmt.Sprintf("%s/assets/%s.jpg", serverURL, hex.EncodeToString([]byte(userinfo.ID))))
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
			row.AddCell().SetValue(bennyEpoch.Format(dateFormat))
		} else {
			row.AddCell().SetValue(userinfo.CreatedAt.Format(dateFormat))
		}
		if userinfo.UpdatedAt.Before(bennyEpoch) {
			row.AddCell().SetValue(bennyEpoch.Format(dateFormat))
		} else {
			row.AddCell().SetValue(userinfo.UpdatedAt.Format(dateFormat))
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
