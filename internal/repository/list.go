package repository

import (
	"GoToDoList/internal/model"
	"gorm.io/gorm"
)

type ListRepository struct {
	db *gorm.DB
}

// NewListRepository 创建ListRepository实例
func NewListRepository(db *gorm.DB) *ListRepository {
	return &ListRepository{db: db}
}

// CreateList 创建列表
func (r *ListRepository) CreateList(list *model.List) (*gorm.DB, *model.List) {
	return r.db.Create(list), list
}

// GetListByID 根据ID获取列表
func (r *ListRepository) GetListByID(id uint) (*model.List, error) {
	var list model.List
	result := r.db.First(&list, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &list, nil
}

// UpdateList 更新列表
func (r *ListRepository) UpdateList(list *model.List) *gorm.DB {
	return r.db.Save(list)
}

// DeleteList 删除列表
// 由于删除没有的数据时不会返回错误，所以这里返回受影响的行数和错误
func (r *ListRepository) DeleteList(id uint) (int64, error) {
	result := r.db.Delete(&model.List{}, id)
	return result.RowsAffected, result.Error
}

// GetUserByName 根据用户名获取用户
func (r *ListRepository) GetUserByName(username string) (*model.User, error) {
	var user model.User
	result := r.db.Where("username=?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil //用户不存在
		}
		return nil, result.Error //可能数据库查询有误
	}
	return &user, nil
}

// GetListByListName 根据列表名获取列表
func (r *ListRepository) GetListByListName(listName string) (*model.List, error) {
	var list model.List
	result := r.db.Where("list_name=?", listName).First(&list)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil //列表不存在
		}
		return nil, result.Error //可能数据库查询有误
	}
	return &list, nil
}
