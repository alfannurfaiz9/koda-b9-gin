package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Response struct {
	Success bool
	Data    any
	Message string
}

type User struct {
	Email    string
	Password string
}

var Users = []User{{
	Email:    "niko",
	Password: "niko123",
}, {
	Email:    "given",
	Password: "given123",
}}

func main() {
	router := gin.Default()

	router.POST("/register", func(ctx *gin.Context) {
		var user User

		if err := ctx.ShouldBindWith(&user, binding.JSON); err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Data:    user,
				Message: "internal server error",
			})

			return
		}

		if user.Email == "" || user.Password == "" {
			ctx.JSON(http.StatusNotFound, Response{
				Success: false,
				Data:    user,
				Message: "body not found",
			})

			return
		}

		if len(user.Email) < 6 || len(user.Password) < 6 {
			ctx.JSON(http.StatusNotAcceptable, Response{
				Success: false,
				Data:    user,
				Message: "email or password must be at least 6 character",
			})

			return
		}

		for _, v := range Users {
			if user.Email == v.Email {
				ctx.JSON(http.StatusNotAcceptable, Response{
					Success: false,
					Data:    user,
					Message: "account already exist",
				})

				return
			}
		}

		ctx.JSON(http.StatusCreated, Response{
			Success: true,
			Data:    user,
			Message: "account created",
		})

		Users = append(Users, user)
	})

	router.POST("/login", func(ctx *gin.Context) {
		var user User

		if err := ctx.ShouldBindWith(&user, binding.JSON); err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Data:    user,
				Message: "internal server error",
			})

			return
		}

		for _, v := range Users {
			if user.Email == v.Email && user.Password == v.Password {
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Data:    user,
					Message: "login success",
				})

				return
			}
		}

		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Data:    user,
			Message: "invalid email or password",
		})
	})

	router.Run("localhost:9000")
}
