package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"learning/restapi/app"
	"learning/restapi/controller"
	"learning/restapi/exception"
	"learning/restapi/helper"
	"learning/restapi/middleware"
	"learning/restapi/repository/category"
	"learning/restapi/repository/user"
	category2 "learning/restapi/service/category"
	user2 "learning/restapi/service/user"
	"net/http"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	//user
	userRepository := user.NewUserRepository()
	userService := user2.NewUserService(userRepository, db, validate)
	userController := controller.NewAuthController(userService)
	//category
	categoryRepository := category.NewCategoryRepository()
	categoryService := category2.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.POST("/api/auth/register", userController.Register)

	router.GET("/api/categories", categoryController.GetAll)
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
