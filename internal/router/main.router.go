package router

import "github.com/gin-gonic/gin"

func MainRouter(router *gin.Engine) {
	initRegisterRouter(router)
	initLoginRouter(router)
}
