package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/handler"
)

func AuthRouter(router *gin.Engine, api handler.Handler) {

	router.POST("/signup", api.SignUp)
	router.POST("/signin", api.SignIn)

}
