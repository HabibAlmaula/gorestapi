package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id         string `gorm:"primaryKey;type:varchar(36)"`
	FullName   string
	Email      string `gorm:"unique;not null;index;type:varchar(50)"`
	Password   string
	Categories []Category
	DeletedAt  gorm.DeletedAt
}
