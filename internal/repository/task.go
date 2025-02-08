package repository

import (
	"GoToDoList/internal/model"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// CreateTask 创建任务
func (r *TaskRepository) CreateTask(task *model.Task) *gorm.DB {
	return r.db.Create(task)
}

// GetTaskByID 根据ID获取任务
func (r *TaskRepository) GetTaskByID(id uint) (*model.Task, error) {
	var task model.Task
	result := r.db.First(&task, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &task, nil
}

// UpdateTask 更新任务
func (r *TaskRepository) UpdateTask(task *model.Task) *gorm.DB {
	// 使用 Select 方法仅更新非零值字段
	result := r.db.Model(&model.Task{}).Select("ListID", "Name", "Description", "Priority", "DueDate", "Completed").Where("id = ?", task.ID).Updates(task)
	return result
}

// DeleteTask 删除任务
func (r *TaskRepository) DeleteTask(id uint) (int64, error) {
	result := r.db.Delete(&model.Task{}, id)
	return result.RowsAffected, result.Error
}
