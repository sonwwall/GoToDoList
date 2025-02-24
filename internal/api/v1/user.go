package v1

import (
	"GoToDoList/internal/model"
	"GoToDoList/internal/service"
	"GoToDoList/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// 用户注册
func (h *UserHandler) UserRegister(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})

		return
	}

	//获取文件
	file, header, err := c.Request.FormFile("avatar")
	if err != nil && err.Error() != "http: no such file" {
		c.JSON(200, gin.H{
			"code":  400,
			"msg":   "上传文件失败",
			"error": err.Error(),
		})
		return
	}

	if err := h.UserService.Register(&user, file, header); err != nil {
		if err == service.UserExisted {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "用户已存在",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"code":  400,
				"msg":   "注册失败",
				"error": err.Error(),
			})
			return
		}

	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

// 用户登录
func (h *UserHandler) UserLogin(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	err, token := h.UserService.Login(&user)
	if err != nil {
		if err == service.UserNotExisted {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "用户不存在",
			})
			return
		}
		if err == service.PasswordError {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "密码错误",
			})
			return
		}
		c.JSON(200, gin.H{
			"code":  400,
			"msg":   "登录失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "登录成功",
		"token": token,
	})

}

// 更新用户信息
func (h *UserHandler) UpdateUserInfo(c *gin.Context) {
	// 获取用户名,验证用户
	username, exists := c.Get("username")

	if !exists {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}
	// 获取用户输入的新用户名和昵称
	newusername := c.PostForm("new_username")
	newnickname := c.PostForm("new_nickname")

	//获取文件
	file, header, err := c.Request.FormFile("avatar")
	if err != nil && err.Error() != "http: no such file" {
		c.JSON(200, gin.H{
			"code":  400,
			"msg":   "上传文件失败",
			"error": err.Error(),
		})
		return
	}

	//更新用户信息
	if err := h.UserService.UpdateUserInfo(utils.AnyToString(username), file, header, newusername, newnickname); err != nil {
		c.JSON(200, gin.H{
			"code":  400,
			"msg":   "更新用户信息失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更新用户信息成功",
	})

}
