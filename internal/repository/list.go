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
func (r *ListRepository) GetListByID(id uint, userid uint) (*model.List, error) {
	var list model.List
	result := r.db.Where("id=?", id).Where("user_id=?", userid).First(&list)
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
func (r *ListRepository) DeleteList(id uint) error {
	result := r.db.Delete(&model.List{}, id)
	return result.Error
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

func (r *ListRepository) SearchList(keyword string, page, size, userid uint) ([]*model.List, int64, error) {
	var lists []*model.List
	var total int64
	//公式 (page - 1) * size 计算当前页的起始记录位置。例如：
	//如果 page = 1，size = 10，则 offset = 0（从第 0 条记录开始）。
	//如果 page = 2，size = 10，则 offset = 10（从第 10 条记录开始）
	offset := (page - 1) * size
	result := r.db.Model(&model.List{}).
		Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Where("user_id=?", userid).
		Count(&total). //Count 方法会自动执行统计查询，并返回统计结果。
		Limit(int(size)).
		Offset(int(offset)).
		Find(&lists) //查询最终的结果，并将结果存储到变量 lists 中

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return lists, total, nil
}
