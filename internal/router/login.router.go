package router

import (
	"github.com/alfannurfaiz9/koda-b9-gin.git/internal/controller"
	"github.com/alfannurfaiz9/koda-b9-gin.git/internal/service"
	"github.com/gin-gonic/gin"
)

func initLoginRouter(router *gin.Engine) {
	r := router.Group("/login")

	ls := service.NewLoginService()
	lc := controller.NewLoginController(ls)

	r.POST("", lc.Login)
}
