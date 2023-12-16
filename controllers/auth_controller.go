package controllers

import (
	"github.com/BerdanAkbulut/task-app-backend/exceptions"
	"github.com/BerdanAkbulut/task-app-backend/requests"
	"github.com/BerdanAkbulut/task-app-backend/services"
	"github.com/labstack/echo/v4"
)

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *authController {
	return &authController{authService: authService}
}

func (ac *authController) HandleRoutes(e *echo.Echo) {
	g := e.Group("/auth")

	g.POST("/authenticate", ac.handleAuthenticate)
	g.POST("/register", ac.handleRegister)
}

func (ac *authController) handleRegister(c echo.Context) error {
	req := new(requests.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return exceptions.ThrowHttpError(err.Error(), 400)
	}

	res, ex := ac.authService.Register(req)

	if ex != nil {
		return exceptions.ThrowHttpError(ex.Error(), 400)
	}
	return c.JSON(201, res)
}

func (ac *authController) handleAuthenticate(c echo.Context) error {
	req := new(requests.AuthenticateRequest)
	if err := c.Bind(req); err != nil {
		return exceptions.ThrowHttpError(err.Error(), 400)
	}

	res, ex := ac.authService.Authenticate(req)

	if ex != nil {
		return exceptions.ThrowHttpError(ex.Error(), 400)
	}
	return c.JSON(200, res)
}
