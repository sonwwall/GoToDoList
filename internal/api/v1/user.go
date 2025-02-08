package v1

import (
	"GoToDoList/internal/model"
	"GoToDoList/internal/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})

		return
	}

	if err := service.Register(&user); err != nil {
		if err == service.Userhasalreadyexisted {
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

func UserLogin(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	err, token := service.Login(&user)
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
