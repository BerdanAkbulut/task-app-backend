package utils

import (
	"log"

	"github.com/BerdanAkbulut/task-app-backend/controllers"
	"github.com/BerdanAkbulut/task-app-backend/exceptions"
	"github.com/BerdanAkbulut/task-app-backend/jwt"
	"github.com/BerdanAkbulut/task-app-backend/pkg"
	"github.com/BerdanAkbulut/task-app-backend/repository"
	"github.com/BerdanAkbulut/task-app-backend/services"
	"github.com/labstack/echo/v4"
)

func RunApp(app *pkg.App) {
	e := echo.New()
	e.Use(jwt.AuthMiddleware)
	
	repository.LoadDBInstance()
	LoadControllers(e)
	exceptions.GlobalExceptionHandler(e)

	err := e.Start(":" + app.Port)
	if err != nil {
		log.Fatal("Error whhile starting app: ", err)
	}
}

func LoadControllers(e *echo.Echo) {
	// Task
	taskRepository := repository.NewTaskRepository()
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewTaskController(taskService)
	taskController.HandleRoutes(e)
	// Auth
	userRepository := repository.NewUserRepository()
	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)
	authController.HandleRoutes(e)

	controllers.UserControllers(e)
}
