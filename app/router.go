package app

import (
	"golang-restful-api/controller"
	"golang-restful-api/exception"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"


)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler
	router.MethodNotAllowed = http.HandlerFunc(exception.NotAllowedError)
	router.NotFound = http.HandlerFunc(exception.NotFoundRouteError)

	return router
	
}