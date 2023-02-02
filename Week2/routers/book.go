package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/database"
	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/handler"
)

func Router() *gin.Engine {
	router := gin.Default() //get the default engine for future customization
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	}
	router.GET("/books", api.GetBooks) //set the function for http://localhost:8080/books : Get request
	//while calling handler function, gin will pass*gin.Context in the handler function
	router.POST("/books", api.SaveBook)
	return router

}
