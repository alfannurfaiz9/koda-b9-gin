package controller

import (
	"net/http"

	"github.com/alfannurfaiz9/koda-b9-gin.git/internal/dto"
	"github.com/alfannurfaiz9/koda-b9-gin.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ILoginService interface {
	ValidateLogin(user dto.User) error
}

type LoginController struct {
	ls ILoginService
}

func NewLoginController(ls ILoginService) *LoginController {
	return &LoginController{
		ls: ls,
	}
}

func (l *LoginController) Login(ctx *gin.Context) {
	var user dto.User

	if err := ctx.ShouldBindWith(&user, binding.JSON); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    user,
			Message: "internal server error",
		})

		return
	}

	if err := l.ls.ValidateLogin(user); err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data:    user,
			Message: "login success",
		})

		return
	}

	ctx.JSON(http.StatusNotFound, dto.Response{
		Success: false,
		Data:    user,
		Message: "invalid email or password",
	})
}

type IRegisterService interface {
	LengthValidation(user dto.User) error
	AlreadyRegistered(user dto.User) error
}

type RegisterController struct {
	rc IRegisterService
}

func NewRegisterController(rc IRegisterService) *RegisterController {
	return &RegisterController{
		rc: rc,
	}
}

func (r *RegisterController) Register(ctx *gin.Context) {
	var user dto.User

	if err := ctx.ShouldBindWith(&user, binding.JSON); err != nil {
		var user dto.User
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    user,
			Message: "internal server error",
		})

		return
	}

	if err := r.rc.LengthValidation(user); err != nil {
		ctx.JSON(http.StatusNotAcceptable, dto.Response{
			Success: false,
			Data:    user,
			Message: err.Error(),
		})

		return
	}

	if err := r.rc.AlreadyRegistered(user); err != nil {
		ctx.JSON(http.StatusNotAcceptable, dto.Response{
			Success: false,
			Data:    user,
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Data:    user,
		Message: "account created",
	})

	service.Users = append(service.Users, user)
}
