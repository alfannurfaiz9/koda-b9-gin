package service

import (
	"errors"

	"github.com/alfannurfaiz9/koda-b9-gin.git/internal/dto"
)

var Users = []dto.User{{
	Email:    "niko",
	Password: "niko123",
}, {
	Email:    "given",
	Password: "given123",
}}

type LoginService struct{}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (l *LoginService) ValidateLogin(user dto.User) error {
	for _, v := range Users {
		if user.Email == v.Email && user.Password == v.Password {
			return errors.New("login success")

		}
	}

	return nil
}

type RegisterService struct{}

func NewRegisterService() *RegisterService {
	return &RegisterService{}
}

func (r *RegisterService) LengthValidation(user dto.User) error {
	if len(user.Email) < 6 || len(user.Password) < 6 {
		return errors.New("email or password must be at least 6 character")
	}

	return nil
}

func (r *RegisterService) AlreadyRegistered(user dto.User) error {
	for _, v := range Users {
		if user.Email == v.Email {
			return errors.New("account already exist")
		}
	}

	return nil
}
