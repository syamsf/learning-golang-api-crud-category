package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"syamsf/learning-golang-api-crud-category/app"
	"syamsf/learning-golang-api-crud-category/controller"
	"syamsf/learning-golang-api-crud-category/exception"
	"syamsf/learning-golang-api-crud-category/helper"
	"syamsf/learning-golang-api-crud-category/middleware"
	"syamsf/learning-golang-api-crud-category/repository"
	"syamsf/learning-golang-api-crud-category/service"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("server is running in development mode port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
