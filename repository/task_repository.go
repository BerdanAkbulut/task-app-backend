package repository

import (
	"github.com/BerdanAkbulut/task-app-backend/entity"
)

type taskRepository struct {
}

type TaskRepository interface {
	Save(task *entity.Task)
	GetAll() []*entity.Task
}

func (t *taskRepository) Save(task *entity.Task) {
	DB().Create(&task)
}
func (t *taskRepository) GetAll() []*entity.Task {
	tasks := []*entity.Task{}
	DB().Find(&tasks)
	return tasks
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
}
