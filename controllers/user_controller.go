package controllers

import "github.com/labstack/echo/v4"

func UserControllers(e *echo.Echo) {
	g := e.Group("/users")

	g.GET("", nil)
	g.GET("/:id", handleGetUser)
	g.POST("", handlePostUser)
	g.PUT("/:id", handlePutUser)
	g.DELETE("/:id", handleDeleteUser)
}

func handleGetAllUsers(c echo.Context) error {
	return nil
}

func handleGetUser(c echo.Context) error {
	return nil
}

func handlePostUser(c echo.Context) error {
	return nil

}

func handlePutUser(c echo.Context) error {
	return nil

}

func handleDeleteUser(c echo.Context) error {
	return nil
}
