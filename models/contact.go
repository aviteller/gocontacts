package models

import (
	"fmt"

	u "github.com/avi/go-contacts/utils"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Please enter contact name"), false
	}
	if contact.Phone == "" {
		return u.Message(false, "Please enter contact number"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "Please login"), false
	}

	return u.Message(true, "success"), true
}

func DeleteContact(id uint) map[string]interface{} {

	GetDB().Table("contacts").Where("id = ?", id).Delete(id)

	res := u.Message(true, "success")

	return res
}

func (contact *Contact) Create() map[string]interface{} {
	if res, ok := contact.Validate(); !ok {
		return res
	}

	GetDB().Create(contact)

	res := u.Message(true, "success")
	res["contact"] = contact
	return res
}

func GetContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}

	return contact
}

func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}
