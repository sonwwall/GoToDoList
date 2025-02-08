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

// 登录
var UserNotExisted = errors.New("用户不存在")
var PasswordError = errors.New("密码错误")

func Login(user *model.User) (error, string) {
	// 创建一个 UserRepository 实例，用于与数据库交互
	repo := repository.NewUserRepository(global.Mysql)
	existingUser, err := repo.GetUserByUsername(user.Username) //查询数据库有无此人,如果存在就返回该用户实例
	// 如果没有发生错误但未找到用户，则返回用户不存在的错误。
	if err == nil && existingUser == nil {
		return UserNotExisted, ""
	}
	// 如果发生错误，直接返回错误。
	if err != nil {
		return err, ""
	}
	// 验证用户提供的密码是否与存储的密码匹配。
	if !auth.CheckPasswordHash(user.Password, existingUser.Password) {
		return PasswordError, ""
	} else {

		token, err := auth.GenerateToken(existingUser.Username)
		if err != nil {
			return err, ""
		}

		return nil, token
	}

}
