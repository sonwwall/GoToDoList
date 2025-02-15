package service

import (
	"GoToDoList/internal/model"
	"GoToDoList/internal/repository"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
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
func (s *ListService) CreateList(list *model.List, file multipart.File, header *multipart.FileHeader) error {

	result, relist := s.repo.CreateList(list)
	if result.Error != nil {
		return result.Error
	}

	// 如果用户上传了图片，则更新图片
	if file != nil && header != nil {
		// 检查文件类型
		ext := filepath.Ext(header.Filename)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			return errors.New("文件类型不支持")
		}

		// 将得到的用户信息中的id传入
		url, err := s.UpdateDescPicture(relist.ID, file, header)
		if err != nil {
			return err
		}
		relist.DescPicture = url

		//再次存入数据库，这次是为了存入头像url
		if err := s.repo.UpdateList(relist); err != nil {
			return err.Error
		}
	}
	return nil

}

// GetListByID 根据给定的 ID 获取列表。
func (s *ListService) GetListByID(id uint, userid uint) (*model.List, error) {
	return s.repo.GetListByID(id, userid)
}

// UpdateList 更新列表。
func (s *ListService) UpdateList(list *model.List, file multipart.File, header *multipart.FileHeader) error {
	// 如果用户上传了图片，则更新图片
	if file != nil && header != nil {
		// 检查文件类型
		ext := filepath.Ext(header.Filename)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			return errors.New("文件类型不支持")
		}

		// 将得到的用户信息中的id传入
		url, err := s.UpdateDescPicture(list.ID, file, header)
		if err != nil {
			return err
		}
		list.DescPicture = url

	}
	result := s.repo.UpdateList(list)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

// DeleteList 根据给定的 ID 删除列表。

var ErrListNotFound = errors.New("列表不存在")

func (s *ListService) DeleteList(id uint) error {
	return s.repo.DeleteList(id)
}

// GetUserByName 根据给定的用户名获取用户。
func (s *ListService) GetUserByName(username string) (*model.User, error) {
	return s.repo.GetUserByName(username)
}

// 更新描述图像
func (s *ListService) UpdateDescPicture(listID uint, file multipart.File, header *multipart.FileHeader) (string, error) {
	// 创建保存路径
	dir := "./uploads/list_desc_pictures"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", fmt.Errorf("创建目录失败:%v", err)
	}

	// 生成文件名
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d_%s%s", listID, time.Now().Format("2006-01-02-15-04-05"), ext)
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
	DescPictureURL := fmt.Sprintf("/uploads/list_desc_pictures/%s", filename)

	return DescPictureURL, nil
}

// SearchList 根据给定的关键词搜索列表。
func (s *ListService) SearchList(keyword string, page, size, userid uint) ([]*model.List, int64, error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}
	return s.repo.SearchList(keyword, page, size, userid)
}

// SearchListAndTasks 根据给定的关键词搜索列表和任务。
func (s *ListService) SearchListAndTasks(keyword string, page, size, userid uint) ([]*model.List, int64, error) {
	if page == 0 {
		page = 1
	}
	size = 1
	return s.repo.SearchListAndTasks(keyword, page, size, userid)
}
