package v1

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/api/v1/user")
	{
		userGroup.POST("/register", UserRegister)
		userGroup.POST("/login", UserLogin)
	}

	return router
}
