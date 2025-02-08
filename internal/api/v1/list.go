package v1

import (
	"GoToDoList/internal/model"
	"GoToDoList/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
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

// GetList 获取列表
func (h *ListHandler) GetList(c *gin.Context) {
	id := c.Param("id")
	// 将id转换为int类型
	listID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	list, err := h.ListService.GetListByID(uint(listID))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取清单失败",
		})
	}
	if list == nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "清单不存在",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取清单成功",
		"data": list,
	})

}

// UpdateList 更新列表
func (h *ListHandler) UpdateList(c *gin.Context) {
	id := c.Param("id")
	// 将id转换为int类型
	listID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	// 先获取要更新的清单
	existinglist, err := h.ListService.GetListByID(uint(listID))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取清单失败",
		})
	}
	if existinglist == nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "清单不存在",
		})
		return
	}
	var list model.List
	if err := c.ShouldBind(&list); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	list.ID = uint(listID)
	list.CreatedAt = existinglist.CreatedAt //如果不加这一条就会使得创建时间变为空
	// 更新列表
	if err := h.ListService.UpdateList(&list); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "更新清单失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更新清单成功",
		"data": list,
	})
}

// DeleteList 删除列表
func (h *ListHandler) DeleteList(c *gin.Context) {
	id := c.Param("id")
	// 将id转换为int类型
	listID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	if err := h.ListService.DeleteList(uint(listID)); err != nil {
		if err == service.ErrListNotFound {
			c.JSON(200, gin.H{
				"code": 404,
				"msg":  "清单不存在",
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "删除清单失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除清单成功",
	})

}
