package repository

import (
	"GoToDoList/internal/model"
	"gorm.io/gorm"
)

type ListRepository struct {
	db *gorm.DB
}

// 创建ListRepository实例
func NewListRepository(db *gorm.DB) *ListRepository {
	return &ListRepository{db: db}
}

// 创建列表
func (r *ListRepository) CreateList(list *model.List) *gorm.DB {
	return r.db.Create(list)
}

// 根据ID获取列表
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

// 更新列表
func (r *ListRepository) UpdateList(list *model.List) *gorm.DB {
	return r.db.Save(list)
}

// 删除列表
// 由于删除没有的数据时不会返回错误，所以这里返回受影响的行数和错误
func (r *ListRepository) DeleteList(id uint) (int64, error) {
	result := r.db.Delete(&model.List{}, id)
	return result.RowsAffected, result.Error
}
