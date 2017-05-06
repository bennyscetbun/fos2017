package main

import "time"

type TShirtSize int

const (
	_              = iota
	TShirtSizeBoyS = TShirtSize(iota)
	TShirtSizeBoyM
	TShirtSizeBoyL
	TShirtSizeBoyXL
	TShirtSizeBoyXXL
	TShirtSizeGirlS
	TShirtSizeGirlM
	TShirtSizeGirlL
	TShirtSizeTankS
	TShirtSizeTankM
	TShirtSizeTankL
	TShirtSizeTankXL
)

type Regime int

const (
	_         = iota
	RegimeAll = Regime(iota)
	RegimeVegetarian
	RegimeVegan
)

type EmergencyContactType int

const (
	_                       = iota
	EmergencyContactTypeDad = EmergencyContactType(iota)
	EmergencyContactTypeMum
	EmergencyContactTypeFamilly
	EmergencyContactTypePartner
	EmergencyContactTypeOther
)

type EmergencyContact struct {
	Firstname            string
	Lastname             string
	Address              string
	CP                   string
	Town                 string
	PhoneNumber          string
	EmergencyContactType EmergencyContactType
}

type EnglishLevel int

const (
	_                = iota
	EnglishLevelNone = EnglishLevel(iota)
	EnglishLevelABit
	EnglishLevelSchool
	EnglishLevelGood
	EnglishLevelFluent
	EnglishLevelBillangual
)

type JobsType int

const (
	_                      = iota
	JobsTypeAcceuilArtiste = JobsType(iota)
	JobsTypeAcceuilPublic
	JobsTypeBackline
	JobsTypeCaisse
	JobsTypeEcocup
	JobsTypeEnvironment
	JobsTypeRestauration
	JobsTypeMerchandising
	JobsTypeMontage
	JobsTypeRuns
	JobsTypeBarman
)

type DayThere int

const (
	DayThereSunday1 = DayThere(1 << iota)
	DayThereMonday1
	DayThereTuesday1
	DayThereWesnesday1
	DayThereThursday1
	DayThereFriday1
	DayThereSaturday1
	DayThereSunday2
	DayThereMonday2
	DayThereTuesday2
	DayThereWesnesday2
)

type BoolOrEmpty uint8

const (
	_               = iota
	BoolOrEmptyTrue = BoolOrEmpty(iota)
	BoolOrEmptyFalse
)

type UserInfo struct {
	ID                    string `gorm:"primary_key"`
	Firstname             string
	Lastname              string
	Address               string
	CP                    string
	Town                  string
	HealthNumber          string
	BirthDate             time.Time `gorm:"type:DATETIME"`
	BirthPlace            string
	PhoneNumber           string
	Facebook              *string
	TShirt                TShirtSize
	Regime                Regime
	Allergy               string
	MedicalInfo           string
	DriverLicenceVL       BoolOrEmpty
	DriverLicencePL       BoolOrEmpty
	FirstAidTraining      BoolOrEmpty
	EnglishLevel          EnglishLevel
	OtherLanguage         string
	AlreadyBeenBenevol    string
	AlreadyBeenBenevolFOS string
	DidYouCameFOS         BoolOrEmpty
	WhatYouWantToDo1      JobsType
	WhatYouWantToDo2      JobsType
	WhatYouWantToDo3      JobsType
	WhatYouWantToDo4      JobsType
	OtherJobs             string
	WhenCanYouBeThere     DayThere
	OtherInfo             string
	Photo                 string

	EmergencyContactFirstname   string
	EmergencyContactLastname    string
	EmergencyContactAddress     string
	EmergencyContactCP          string
	EmergencyContactTown        string
	EmergencyContactPhoneNumber string
	EmergencyContactType        EmergencyContactType
}

type User struct {
	ID   string `gorm:"primary_key"`
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	// Auth
	Email    string
	Password string

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Recover
	RecoverToken       string
	RecoverTokenExpiry time.Time

	// Remember is in another table
	IsAdmin bool
}

type Token struct {
	TokenKey  string `gorm:"not null;index"`
	Value     string
	CreatedAt time.Time
}
