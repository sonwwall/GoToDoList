package service

import (
	"GoToDoList/internal/model"
	"GoToDoList/internal/repository"
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

func (s *ListService) CreateList(list *model.List) error {
	return s.repo.CreateList(list).Error
}
