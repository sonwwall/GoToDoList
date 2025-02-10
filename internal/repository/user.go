package repository

import (
	"GoToDoList/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建一个新的UserRepository实例，接收一个数据库连接对象作为参数，并返回该实例。
// 该函数是一个工厂函数，用于创建一个新的UserRepository实例。
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// / CreateUser 创建用户
func (r *UserRepository) CreateUser(user *model.User) *gorm.DB {
	return r.db.Create(user)
}

// / GetUserByUsername 根据用户名获取用户
func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.db.Where("username=?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil //用户不存在
		}
		return nil, result.Error //可能数据库查询有误
	}
	return &user, nil //用户存在

}

// UpdateUser 更新用户资料
func (r *UserRepository) UpdateUser(user *model.User) error {
	return r.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
}
