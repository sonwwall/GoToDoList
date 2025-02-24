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

type ListSearch struct {
	Keyword string `form:"keyword" json:"keyword" binding:"required"`
	Page    uint   `form:"page" json:"page"`
	Size    uint   `form:"size" json:"size"`
	Tag     string `form:"tag" json:"tag"`
}

type ListSearchByGroup struct {
	Page    uint `form:"page" json:"page"`
	Size    uint `form:"size" json:"size"`
	GroupID uint `form:"group_id" json:"group_id"`
}

type ListSearchByTag struct {
	Page uint   `form:"page" json:"page"`
	Size uint   `form:"size" json:"size"`
	Tag  string `form:"tag" json:"tag"`
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
	//获取文件
	file, header, err := c.Request.FormFile("desc_picture")
	if err != nil && err.Error() != "http: no such file" {
		c.JSON(200, gin.H{
			"code":  400,
			"msg":   "上传文件失败",
			"error": err.Error(),
		})
		return
	}
	if err := h.ListService.CreateList(&list, file, header); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "创建失败",
		})
		return
	}
	username_any, _ := c.Get("username")
	username := username_any.(string)

	user, _ := h.ListService.GetUserByName(username)
	list.UserID = user.ID
	if err := h.ListService.UpdateList(&list, file, header); err != nil {
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
	useridany, ok := c.Get("userid")
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "token解析出现问题，请检查",
		})
		return
	}

	userid, _ := useridany.(uint)

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

	list, err := h.ListService.GetListByID(uint(listID), userid)

	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取清单失败",
		})
		return
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
	// 获取用户id,验证用户
	useridany, ok := c.Get("userid")
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "token解析出现问题，请检查",
		})
		return
	}

	userid, _ := useridany.(uint)

	//获取文件
	file, header, err := c.Request.FormFile("desc_picture")
	if err != nil && err.Error() != "http: no such file" {
		c.JSON(200, gin.H{
			"code":  400,
			"msg":   "上传文件失败",
			"error": err.Error(),
		})
		return
	}

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
	existinglist, err := h.ListService.GetListByID(uint(listID), userid)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取清单失败",
		})
		return
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
	list.UserID = existinglist.UserID
	list.ID = uint(listID)
	list.CreatedAt = existinglist.CreatedAt //如果不加这一条就会使得创建时间变为空
	// 更新列表
	if err := h.ListService.UpdateList(&list, file, header); err != nil {
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
	useridany, ok := c.Get("userid")
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "token解析出现问题，请检查",
		})
		return
	}

	userid, _ := useridany.(uint)

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
	//复用获取清单的代码
	existinglist, err := h.ListService.GetListByID(uint(listID), userid)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取清单失败",
		})
		return
	}
	if existinglist == nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "清单不存在",
		})
		return
	}

	if err := h.ListService.DeleteList(uint(listID)); err != nil {

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

// SearchList 搜索列表
func (h *ListHandler) SearchList(c *gin.Context) {
	var search ListSearch
	if err := c.ShouldBind(&search); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	useridany, ok := c.Get("userid")
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "token解析出现问题，请检查",
		})
		return
	}

	userid, _ := useridany.(uint)
	lists, total, err := h.ListService.SearchList(search.Keyword, search.Page, search.Size, userid)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "搜索失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "搜索成功",
		"data": gin.H{
			"lists": lists,
			"total": total,
		},
	})

}

// SearchListAndTasks 搜索列表和任务
func (h *ListHandler) SearchListAndTasks(c *gin.Context) {
	var search ListSearch
	if err := c.ShouldBind(&search); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	useridany, ok := c.Get("userid")
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "token解析出现问题，请检查",
		})
		return
	}

	userid, _ := useridany.(uint)
	lists, total, err := h.ListService.SearchListAndTasks(search.Keyword, search.Page, search.Size, userid)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  400,
			"msg":   "搜索失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "搜索成功",
		"data": gin.H{
			"lists": lists,
			"total": total,
		},
	})

}

// SearchListByGroup 根据组别搜索列表
func (h *ListHandler) SearchListByGroup(c *gin.Context) {
	var search ListSearchByGroup
	if err := c.ShouldBind(&search); err != nil {
		c.JSON(200, gin.H{
			"code":  400,
			"msg":   "参数错误",
			"error": err.Error(),
		})
		return
	}
	useridany, ok := c.Get("userid")
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "token解析出现问题，请检查",
		})
		return
	}

	userid, _ := useridany.(uint)
	lists, total, err := h.ListService.SearchListByGroup(search.GroupID, search.Page, search.Size, userid)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  400,
			"msg":   "搜索失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "搜索成功",
		"data": gin.H{
			"lists": lists,
			"total": total,
		},
	})

}

// SearchListByTag 根据Tag搜索列表
func (h *ListHandler) SearchListByTag(c *gin.Context) {
	var search ListSearchByTag
	if err := c.ShouldBind(&search); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	useridany, ok := c.Get("userid")
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "token解析出现问题，请检查",
		})
		return
	}

	userid, _ := useridany.(uint)
	lists, total, err := h.ListService.SearchListByTag(search.Tag, search.Page, search.Size, userid)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  400,
			"msg":   "搜索失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "搜索成功",
		"data": gin.H{
			"lists": lists,
			"total": total,
		},
	})
}
