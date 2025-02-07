package service

import (
	"GoToDoList/internal/global"
	"GoToDoList/internal/model"
	"GoToDoList/internal/pkg/auth"
	"GoToDoList/internal/repository"
	"errors"
)

// 用户已存在
var Userhasalreadyexisted = errors.New("用户已存在")

// 注册
func Register(user *model.User) error {
	repo := repository.NewUserRepository(global.Mysql)
	existingUser, err := repo.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return Userhasalreadyexisted
	}

	// 对密码进行哈希处理
	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result := repo.CreateUser(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
