package v1

import (
	"GoToDoList/internal/model"
	"GoToDoList/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TaskHandler struct {
	TaskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
	}
}

// CreateTask 创建任务
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	if err := h.TaskService.CreateTask(&task); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "创建任务失败",
		})
		return
	}

	username_any, _ := c.Get("username")
	username := username_any.(string)

	user, _ := h.TaskService.GetUserByName(username)
	task.UserID = user.ID
	fmt.Println(task.UserID)
	if err := h.TaskService.UpdateTask(&task); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "创建失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建任务成功",
	})
}

// GetTask 获取任务
func (h *TaskHandler) GetTask(c *gin.Context) {
	// 获取用户名,验证用户
	username_any, _ := c.Get("username")
	username := username_any.(string)
	user, _ := h.TaskService.GetUserByName(username)

	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	task, err := h.TaskService.GetTask(uint(taskID))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取任务失败",
		})
		return
	}
	if task == nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "任务不存在",
		})
		return
	}
	if task.UserID != user.ID {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "任务不存在(无权限)",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取任务成功",
		"data": task,
	})
}

// UpdateTask 更新任务
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	// 获取用户名,验证用户
	username_any, _ := c.Get("username")
	username := username_any.(string)
	user, _ := h.TaskService.GetUserByName(username)

	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	existingtask, err := h.TaskService.GetTask(uint(taskID))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取任务失败",
		})
		return
	}
	if existingtask == nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "任务不存在",
		})
		return
	}
	// 验证用户
	if existingtask.UserID != user.ID {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "任务不存在(无权限)",
		})
		return
	}
	var task model.Task
	err = c.ShouldBind(&task)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	task.UserID = existingtask.UserID
	task.ID = uint(taskID)
	task.CreatedAt = existingtask.CreatedAt
	if err := h.TaskService.UpdateTask(&task); err != nil {
		if err == service.ErrTaskNotFound {
			c.JSON(200, gin.H{
				"code": 404,
				"msg":  "任务不存在",
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "更新任务失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更新任务成功",
		"data": task,
	})
}

// DeleteTask 删除任务
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	// 获取用户名,验证用户
	username_any, _ := c.Get("username")
	username := username_any.(string)
	user, _ := h.TaskService.GetUserByName(username)

	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	existingtask, err := h.TaskService.GetTask(uint(taskID))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取清单失败",
		})
	}
	if existingtask == nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "清单不存在",
		})
		return
	}
	// 验证用户
	if existingtask.UserID != user.ID {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "任务不存在(无权限)",
		})
		return
	}
	if err := h.TaskService.DeleteTask(uint(taskID)); err != nil {
		if err == service.ErrTaskNotFound {
			c.JSON(200, gin.H{
				"code": 404,
				"msg":  "任务不存在",
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "删除任务失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除任务成功",
	})
}
