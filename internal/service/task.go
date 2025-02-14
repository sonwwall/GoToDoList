package service

import (
	"GoToDoList/internal/model"
	"GoToDoList/internal/repository"
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
func (s *TaskService) GetTask(id uint, userid uint) (*model.Task, error) {
	return s.repo.GetTaskByID(id, userid)
}

// UpdateTask 更新任务
func (s *TaskService) UpdateTask(task *model.Task) error {
	return s.repo.UpdateTask(task).Error
}

// DeleteTask 删除任务
func (s *TaskService) DeleteTask(id uint) error {
	err := s.repo.DeleteTask(id)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByName 根据用户名获取用户
func (s *TaskService) GetUserByName(username string) (*model.User, error) {
	return s.repo.GetUserByName(username)
}

// SearchTask 搜索任务
func (s *TaskService) SearchTask(keyword string, page, size, userid uint) ([]*model.Task, int64, error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}
	return s.repo.SearchTask(keyword, page, size, userid)
}
