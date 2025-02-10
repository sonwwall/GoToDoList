package service

import (
	"GoToDoList/internal/global"
	"GoToDoList/internal/model"
	"GoToDoList/internal/pkg/auth"
	"GoToDoList/internal/repository"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
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

// 更新头像
func UpdateAvatar(userID uint, file multipart.File, header *multipart.FileHeader) (string, error) {
	// 创建保存路径
	dir := "./uploads/avatars"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", fmt.Errorf("创建目录失败:%v", err)
	}

	// 生成文件名
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d_%s%s", userID, time.Now().Format("2006-01-02-15-04-05"), ext)
	// 使用 filepath.Join 将目录路径和文件名组合成完整的文件路径。
	filePath := filepath.Join(dir, filename)

	// 保存文件
	// 使用 os.Create 创建一个新的文件，路径为 filePath
	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("创建文件失败:%v", err)
	}
	// 使用 defer out.Close() 确保文件在函数结束时被关闭。
	defer out.Close()

	// 使用 io.Copy 将上传的文件内容复制到新创建的文件中
	_, err = io.Copy(out, file)
	if err != nil {
		return "", fmt.Errorf("复制文件失败:%v", err)
	}

	// 返回保存的URL
	avatarURL := fmt.Sprintf("/uploads/avatars/%s", filename)

	return avatarURL, nil

}

// 更新用户信息
func UpdateUserInfo(username string, file multipart.File, header *multipart.FileHeader, newusername, newnickname string) error {
	repo := repository.NewUserRepository(global.Mysql)

	//获取已存在的用户信息
	existingUser, err := repo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	if file != nil && header != nil {
		// 检查文件类型
		ext := filepath.Ext(header.Filename)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			return errors.New("文件类型不支持")
		}

		// 更新用户头像
		// 将得到的用户信息中的id传入
		avatarurl, err := UpdateAvatar(existingUser.ID, file, header)
		if err != nil {
			return err
		}
		existingUser.Avatar = avatarurl
	}

	//更新用户信息
	existingUser.Username = newusername
	existingUser.Nickname = newnickname

	if err := repo.UpdateUser(existingUser); err != nil {
		return err
	}

	//获取请求上下文的jwt

	authHeader := global.GinContext.GetHeader("Authorization")
	tokenstring := strings.Replace(authHeader, "Bearer ", "", 1)

	if err := auth.AddTokenToBlacklist(tokenstring); err != nil {
		return err
	}
	return nil

}
