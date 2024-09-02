package app

import (
	"database/sql"
	"fmt"
	"learning/restapi/configs"
	"learning/restapi/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open(configs.Config.DB.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.Config.DB.User, configs.Config.DB.Password, configs.Config.DB.Host, configs.Config.DB.Port, configs.Config.DB.Name))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(configs.Config.DB.MaxOpenConns)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
