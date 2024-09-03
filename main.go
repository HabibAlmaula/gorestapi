package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"learning/restapi/app"
	"learning/restapi/controller"
	"learning/restapi/helper"
	"learning/restapi/repository/category"
	"learning/restapi/repository/user"
	category2 "learning/restapi/service/category"
	user2 "learning/restapi/service/user"
	"net/http"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	// Initialize user-related components
	userRepository := user.NewUserRepository(db)
	userService := user2.NewUserService(userRepository, db, validate)
	userController := controller.NewAuthController(userService)

	// Initialize category-related components
	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category2.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(userController, categoryController)
	// Configure and start the server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	// Start listening for requests
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
