package controllers

import (
	"github.com/BerdanAkbulut/task-app-backend/entity"
	"github.com/BerdanAkbulut/task-app-backend/exceptions"
	"github.com/BerdanAkbulut/task-app-backend/services"
	"github.com/labstack/echo/v4"
)

type taskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) *taskController {
	return &taskController{taskService: taskService}
}

func (tc *taskController) HandleRoutes(e *echo.Echo) {
	g := e.Group("/tasks")
	g.GET("", tc.handleGetAllTasks)
	g.GET("/:id", tc.handleGetTask)
	g.POST("", tc.handlePostTask)
	g.PUT("/:id", tc.handlePutTask)
	g.DELETE("/:id", tc.handleDeleteTask)
}

func (tc *taskController) handleGetAllTasks(c echo.Context) error {
	return c.JSON(200, tc.taskService.GetAll())
}

func (tc *taskController) handleGetTask(c echo.Context) error {
	return nil
}

func (tc *taskController) handlePostTask(c echo.Context) error {

	task := new(entity.Task)
	if err := c.Bind(task); err != nil {
		return exceptions.ThrowHttpError("Unable to bind request body to task", 400)
	}
	tc.taskService.Save(task)
	return c.NoContent(201)
}

func (tc *taskController) handlePutTask(c echo.Context) error {
	return nil

}

func (tc *taskController) handleDeleteTask(c echo.Context) error {
	return nil
}
