package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name     string
	Lastname string
	Dni      string
	Address  string
	Phone    string
	Email    string
	Users    []UsersDevs `gorm:"many2many:user_contacts;foreignKey:ID;joinForeignKey:FkContact;References:ID;joinReferences:FkUser"`
}
