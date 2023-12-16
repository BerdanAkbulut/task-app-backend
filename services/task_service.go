package services

import (
	"github.com/BerdanAkbulut/task-app-backend/entity"
	"github.com/BerdanAkbulut/task-app-backend/repository"
)

type taskService struct {
	taskRepository repository.TaskRepository
}

type TaskService interface {
	Save(task *entity.Task)
	GetAll() []*entity.Task
}

func (service *taskService) Save(task *entity.Task) {
	service.taskRepository.Save(task)
}
func (service *taskService) GetAll() []*entity.Task {
	return service.taskRepository.GetAll()
}

func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	return &taskService{taskRepository: taskRepository}
}
