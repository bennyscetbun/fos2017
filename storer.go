package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/authboss.v1"
)

type DatabaseStorer struct {
	db *gorm.DB
}

func NewDataBaseStorer(db *gorm.DB) *DatabaseStorer {
	//db.DropTable(&User{}, &Token{})
	//db.LogMode(true)
	db.AutoMigrate(&User{}, &Token{}, &UserInfo{})
	return &DatabaseStorer{
		db: db,
	}
}

func (s DatabaseStorer) Create(key string, attr authboss.Attributes) error {
	var user User
	user.RecoverTokenExpiry = time.Now()
	if err := attr.Bind(&user, true); err != nil {
		return err
	}
	user.ID = key
	if err := s.db.Create(&user).Error; err != nil {
		return err
	}
	fmt.Println("Create")
	spew.Dump(user)
	return nil
}

func (s DatabaseStorer) Put(key string, attr authboss.Attributes) error {
	var user User
	if err := s.db.First(&user, "id = ?", key).Error; err != nil {
		return err
	}
	if err := attr.Bind(&user, true); err != nil {
		return err
	}
	user.ID = key
	if user.RecoverTokenExpiry.Before(time.Now()) {
		user.RecoverTokenExpiry = time.Now()
	}
	if err := s.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s DatabaseStorer) Get(key string) (result interface{}, err error) {
	var user User
	if s.db.First(&user, "id = ?", key).RecordNotFound() {
		return nil, authboss.ErrUserNotFound
	}
	return &user, nil
}

func (s DatabaseStorer) PutOAuth(uid, provider string, attr authboss.Attributes) error {
	return s.Put(uid+provider, attr)
}

func (s DatabaseStorer) GetOAuth(uid, provider string) (result interface{}, err error) {
	return s.Get(uid + provider)
}

func (s DatabaseStorer) AddToken(key, value string) error {
	token := Token{
		TokenKey: key,
		Value:    value,
	}
	if err := s.db.Create(&token).Error; err != nil {
		return err
	}
	fmt.Println("AddToken")
	spew.Dump(token)
	return nil
}

func (s DatabaseStorer) DelTokens(key string) error {
	if err := s.db.Where("token_key = ?", key).Delete(&Token{}).Error; err != nil {
		return err
	}
	fmt.Println("DelTokens")
	return nil
}

func (s DatabaseStorer) UseToken(givenKey, token string) error {
	tmpdb := s.db.Delete(&Token{}, "token_key = ? AND value = ?", givenKey, token)
	if tmpdb.RecordNotFound() {
		return authboss.ErrUserNotFound
	} else if tmpdb.Error != nil {
		return tmpdb.Error
	}
	return nil
}

func (s DatabaseStorer) ConfirmUser(tok string) (result interface{}, err error) {
	fmt.Println("==============", tok)

	var user User
	tmpdb := s.db.First(&user, "confirm_token = ?", tok)
	if tmpdb.RecordNotFound() {
		return nil, authboss.ErrUserNotFound
	} else if tmpdb.Error != nil {
		return nil, tmpdb.Error
	}

	return &user, nil
}

func (s DatabaseStorer) RecoverUser(rec string) (result interface{}, err error) {
	var user User
	tmpdb := s.db.First(&user, "recover_token = ?", rec)
	if tmpdb.RecordNotFound() {
		return nil, authboss.ErrUserNotFound
	} else if tmpdb.Error != nil {
		return nil, tmpdb.Error
	}
	//HACK TO LET USER RESEND emails
	if user.Confirmed == false {
		user.Confirmed = true
	}

	return &user, nil
}
