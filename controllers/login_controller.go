package controllers

import (
	"github.com/labstack/echo/v4"
)

type LoginController struct {
}

type UserController struct {
}

func NewUserController() UserController {
	return UserController{}
}

func (us UserController) Login(c echo.Context) error {
	return nil
}
