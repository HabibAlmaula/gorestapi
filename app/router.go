package app

import (
	"github.com/julienschmidt/httprouter"
	"learning/restapi/controller"
	"learning/restapi/exception"
	"learning/restapi/middleware"
)

func NewRouter(userController controller.AuthController, categoryController controller.CategoryController) *httprouter.Router {

	router := httprouter.New()

	// User registration route
	router.POST("/api/auth/register", userController.Register)
	// User login route
	router.POST("/api/auth/login", userController.Login)

	// Category routes with middleware
	router.GET("/api/categories", middleware.NewAuthMiddleware(categoryController.GetAll).ServeHTTP)
	router.GET("/api/categories/:categoryId", middleware.NewAuthMiddleware(categoryController.GetById).ServeHTTP)
	router.POST("/api/categories", middleware.NewAuthMiddleware(categoryController.Create).ServeHTTP)
	router.PUT("/api/categories/:categoryId", middleware.NewAuthMiddleware(categoryController.Update).ServeHTTP)
	router.DELETE("/api/categories/:categoryId", middleware.NewAuthMiddleware(categoryController.Delete).ServeHTTP)
	router.PanicHandler = exception.ErrorHandler

	return router
}
