package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id        int `gorm:"primaryKey;autoIncrement"`
	Name      string
	UserId    string `gorm:"not null;type:varchar(36)"` // Specify the length for VARCHAR
	User      User   `gorm:"foreignKey:UserId;references:Id"`
	DeletedAt gorm.DeletedAt
}
