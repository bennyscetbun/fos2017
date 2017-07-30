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

// wordWrap text at the specified column lineWidth on word breaks

func wordWrap(text string, lineWidth int) []string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return []string{text}
	}
	ret := []string{}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			ret = append(ret, wrapped)
			wrapped = word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	ret = append(ret, wrapped)
	return ret

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
		str := ""
		if len(title) > 0 {
			str = tr(fmt.Sprintf("%s: %s", title, text))
		} else {
			str = tr(fmt.Sprintf("%s", text))
		}
		for _, line := range wordWrap(str, 160) {
			pdf.Text(10, y, line)
			y += fontHeight + 1
		}

		return y + 1
	}

	header := func() float64 {
		const imageSize = 50
		pdf.AddPage()
		_, photoPath := localPhotoPath(userInfo.ID)
		pdf.RegisterImageOptions(photoPath, gofpdf.ImageOptions{})
		imgInfo := pdf.GetImageInfo(photoPath)
		if imgInfo.Height() > imgInfo.Width() {
			pdf.Image(photoPath, 10, 6, 0, imageSize, false, "", 0, "")
		} else {
			pdf.Image(photoPath, 10, 6, imageSize, 0, false, "", 0, "")
		}

		pdf.SetFont("Arial", "B", 15)

		y := (imageSize + 6) * 0.5
		y = centerTextFct(fmt.Sprintf("%s %s", userInfo.Lastname, userInfo.Firstname), y, imageSize+10)
		pdf.SetFont("Arial", "", 12)
		y += 3
		y = centerTextFct(fmt.Sprintf("Tel: %s", userInfo.PhoneNumber), y, imageSize+10)
		return 10 + imageSize + 4
	}

	presence := func(y float64) float64 {
		pdf.SetFont("Arial", "B", 8)
		y = writeUnderFct("Presence", "", y)
		pdf.SetFont("Arial", "", 7)
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
				y += (fontHeight+2)*2 + 1
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
	pdf.SetFont("Arial", "B", 8)
	y = writeUnderFct("Etat Civil", "", y)
	pdf.SetFont("Arial", "", 7)
	y = writeUnderFct("Date/lieu de naissance", fmt.Sprintf("%s/%s", userInfo.BirthDate.Format("02-01-2006"), userInfo.BirthPlace), y)
	y = writeUnderFct("Numero Secu", userInfo.HealthNumber, y)
	y = writeUnderFct("Adresse", fmt.Sprintf("%s %s %s", userInfo.Address, userInfo.CP, userInfo.Town), y)
	y = writeUnderFct("Email", userInfo.ID, y)
	y = writeUnderFct("Tshirt", userInfo.TShirt, y)
	y += 5
	pdf.SetFont("Arial", "B", 8)
	y = writeUnderFct("Compétence", "", y)
	pdf.SetFont("Arial", "", 7)
	y = writeUnderFct("", fmt.Sprintf("Permis VL: %s / Permis PL: %s / Premier Secours: %s / Anglais: %s", userInfo.DriverLicenceVL, userInfo.DriverLicencePL, userInfo.FirstAidTraining, userInfo.EnglishLevel), y)
	y = writeUnderFct("Autres langues parlées", userInfo.OtherLanguage, y)

	y += 5
	y = presence(y)
	y += 5
	pdf.SetFont("Arial", "B", 8)
	y = writeUnderFct("Contact d'urgence", "", y)
	pdf.SetFont("Arial", "", 7)
	y = writeUnderFct("Lien", userInfo.EmergencyContactType, y)
	y = writeUnderFct("Nom", userInfo.EmergencyContactLastname, y)
	y = writeUnderFct("Prénom", userInfo.EmergencyContactFirstname, y)
	y = writeUnderFct("Tel", userInfo.EmergencyContactPhoneNumber, y)
	y = writeUnderFct("Adresse", fmt.Sprintf("%s %s %s", userInfo.EmergencyContactAddress, userInfo.EmergencyContactCP, userInfo.EmergencyContactTown), y)
	y += 5
	pdf.SetFont("Arial", "B", 8)
	y = writeUnderFct("Santé", "", y)
	pdf.SetFont("Arial", "", 7)
	y = writeUnderFct("Information médicales", userInfo.MedicalInfo, y)
	y = writeUnderFct("Allergies", userInfo.Allergy, y)
	y = writeUnderFct("Régime alimentaire", userInfo.Regime, y)
	y += 5
	pdf.SetFont("Arial", "B", 8)
	y = writeUnderFct("Autres infos", "", y)
	pdf.SetFont("Arial", "", 7)
	y = writeUnderFct("Choix Job", fmt.Sprintf("%s,%s,%s,%s", userInfo.WhatYouWantToDo1, userInfo.WhatYouWantToDo2, userInfo.WhatYouWantToDo3, userInfo.WhatYouWantToDo4), y)

	y = writeUnderFct("Choix info", userInfo.OtherJobs, y)

	y = writeUnderFct("Déjà venu au FOS", userInfo.DidYouCameFOS, y)
	y = writeUnderFct("Déjà Bénévole au FOS", userInfo.AlreadyBeenBenevolFOS, y)
	y = writeUnderFct("Déjà Bénévole", userInfo.AlreadyBeenBenevol, y)

	y = writeUnderFct("Info supplementaire", userInfo.OtherInfo, y)
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
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s_%s.pdf\"", userInfo.Lastname, userInfo.Firstname))
	w.Write(buf.Bytes())
	//err := pdf.OutputFileAndClose("hello.pdf")
}
