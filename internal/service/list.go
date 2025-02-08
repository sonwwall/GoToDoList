package service

import (
	"GoToDoList/internal/model"
	"GoToDoList/internal/repository"
	"errors"
)

type ListService struct {
	// repo 是 ListService 使用的仓库实例，用于执行数据相关的操作。
	repo *repository.ListRepository
}

// NewListService 创建并返回一个新的 ListService 实例。
// 参数 repo 是一个 ListRepository 接口，用于定义如何与数据存储进行交互。
// 返回值是 ListService 的实例，用于执行列表相关的操作。
func NewListService(repo *repository.ListRepository) *ListService {
	return &ListService{repo: repo}
}

// CreateList 创建一个新的列表。
func (s *ListService) CreateList(list *model.List) error {
	return s.repo.CreateList(list).Error
}

// GetListByID 根据给定的 ID 获取列表。
func (s *ListService) GetListByID(id uint) (*model.List, error) {
	return s.repo.GetListByID(id)
}

// UpdateList 更新列表。
func (s *ListService) UpdateList(list *model.List) error {
	return s.repo.UpdateList(list).Error
}

// DeleteList 根据给定的 ID 删除列表。

var ErrListNotFound = errors.New("列表不存在")

func (s *ListService) DeleteList(id uint) error {
	rowsAffected, err := s.repo.DeleteList(id)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrListNotFound
	}
	return nil
}
