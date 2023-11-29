package models

import (
	"gorm.io/gorm"
)

type UsersDevs struct {
	gorm.Model
	Email      string    `gorm:"column:email"`
	Password   string    `gorm:"column:password"`
	FkCompany  int       `gorm:"column:fk_company"`
	FkCustomer int       `gorm:"column:fk_customer"`
	FkRole     int       `gorm:"column:fk_role"`
	Company    Company   `gorm:"foreignKey:fk_company"`
	Customer   Customer  `gorm:"foreignKey:fk_customer"`
	Role       Role      `gorm:"foreignKey:fk_role"`
	Contacts   []Contact `gorm:"many2many:user_contacts;foreignKey:ID;joinForeignKey:FkUserDev;References:ID;joinReferences:FkContact"`
	DeletedBy  *uint     `gorm:"column:deleted_by"`
	UpdatedBy  *uint     `gorm:"column:updated_by"`
	State      int       `gorm:"column:state"`
}
