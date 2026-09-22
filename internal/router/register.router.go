package router

import (
	"github.com/alfannurfaiz9/koda-b9-gin.git/internal/controller"
	"github.com/alfannurfaiz9/koda-b9-gin.git/internal/service"
	"github.com/gin-gonic/gin"
)

func initRegisterRouter(router *gin.Engine) {
	r := router.Group("/register")

	rs := service.NewRegisterService()
	rc := controller.NewRegisterController(rs)

	r.POST("", rc.Register)
}
