package v1

import (
	"GoToDoList/internal/model"
	"GoToDoList/internal/service"
	"github.com/gin-gonic/gin"
)

// ListHandler 列表处理器
type ListHandler struct {
	ListService *service.ListService
}

// NewListHandler 创建列表处理器即返回实例
func NewListHandler(listService *service.ListService) *ListHandler {
	return &ListHandler{
		ListService: listService,
	}
}

// CreateList 创建列表
func (h *ListHandler) CreateList(c *gin.Context) {
	var list model.List
	if err := c.ShouldBind(&list); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})

		return
	}
	if err := h.ListService.CreateList(&list); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "创建失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
	})
}
