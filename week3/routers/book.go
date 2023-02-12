package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/handler"
)

func BookRouter(router *gin.Engine, api handler.Handler) {
	// router := gin.Default() //get the default engine for future customization
	// api := handler.Handler{
	// 	DB: database.GetDB(), //set the handler DB
	// }
	router.GET("/books", api.GetBooks) //set the function for http://localhost:8080/books : Get request
	//while calling handler function, gin will pass*gin.Context in the handler function
	router.POST("/books", api.SaveBook)
	router.Get("/books/:id", api.GetBookByID)
	router.DELETE("/books/:id", api.DeleteBookByID)
	router.PUT("/books", api.UpdateBookByID)
	//return router

}

/*

http://localhost:8080/book/11			- param
http://localhost:8080/book?id=11		- query



*/
