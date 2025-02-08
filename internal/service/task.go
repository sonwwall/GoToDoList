package service

import (
	"GoToDoList/internal/model"
	"GoToDoList/internal/repository"
	"errors"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask 创建任务
func (s *TaskService) CreateTask(task *model.Task) error {
	return s.repo.CreateTask(task).Error
}

// GetTask 根据ID获取任务
func (s *TaskService) GetTask(id uint) (*model.Task, error) {
	return s.repo.GetTaskByID(id)
}

// UpdateTask 更新任务
func (s *TaskService) UpdateTask(task *model.Task) error {
	return s.repo.UpdateTask(task).Error
}

// DeleteTask 删除任务
var ErrTaskNotFound = errors.New("任务不存在")

func (s *TaskService) DeleteTask(id uint) error {
	rowsAffected, err := s.repo.DeleteTask(id)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrTaskNotFound
	}
	return nil
}
