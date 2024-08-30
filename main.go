package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"learning/restapi/app"
	"learning/restapi/controller"
	"learning/restapi/exception"
	"learning/restapi/helper"
	"learning/restapi/middleware"
	"learning/restapi/model/base"
	"learning/restapi/model/web/response"
	"learning/restapi/repository"
	"learning/restapi/service"
	"net/http"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:MYPassword@123@tcp(127.0.0.1:3306)/go_restapi")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func categories(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}(db)

	rows, err := db.Query("select id, name from category")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}(rows)

	var result []response.CategoryResponse

	for rows.Next() {
		var each = response.CategoryResponse{}
		var err = rows.Scan(&each.Id, &each.Name)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	res := base.BaseResponse{
		Code:    200,
		Message: "Success Get Category By Id",
		Succes:  true,
		Data:    result,
	}
	//convert result to json
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

}

func main() {
	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.GetAll)
	//router.GET("/api/categories", categories)
	router.GET("/api/categories/:categoryId", categoryController.GetById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
