package app

import (
	"database/sql"
	"fmt"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learning/restapi/configs"
	"learning/restapi/helper"
	"learning/restapi/model/domain"
	"log"
	"time"
)

func NewDB() *gorm.DB {
	db, err := sql.Open(configs.Config.DB.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.Config.DB.User, configs.Config.DB.Password, configs.Config.DB.Host, configs.Config.DB.Port, configs.Config.DB.Name))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(configs.Config.DB.MaxOpenConns)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	gormDB, err := gorm.Open(gormMysql.New(gormMysql.Config{
		Conn: db,
	}), &gorm.Config{})

	// Apply migrations
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	err = gormDB.AutoMigrate(&domain.User{}, &domain.Category{})

	return gormDB
}
