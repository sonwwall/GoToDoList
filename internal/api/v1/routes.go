package v1

import (
	"GoToDoList/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/api/v1/user")
	{
		userGroup.POST("/register", UserRegister)
		userGroup.POST("/login", UserLogin)
	}

	taskGroup := router.Group("/api/v1/task")
	taskGroup.Use(middleware.JwtAuthMiddleware())
	{
		taskGroup.POST("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "test OK",
			})
		})

	}

	return router
}
