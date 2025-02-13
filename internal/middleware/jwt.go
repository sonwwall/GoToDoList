package middleware

import (
	"GoToDoList/internal/global"
	"GoToDoList/internal/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// jwt中间件
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的Authorization字段
		authHeader := c.GetHeader("Authorization")
		// 如果Authorization字段为空，返回401错误信息，并终止请求处理
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请求未携带token",
			})
			c.Abort()
			return
		}

		// 从Authorization字段中提取token字符串，移除"Bearer "前缀
		tokenstring := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := auth.ParseToken(tokenstring)
		if err != nil {
			if err.Error() == "token 已失效" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  "token已失效,请重新登录",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token错误",
			})
			c.Abort()
			return
		}

		// 将解析出的用户名设置到请求上下文中，供后续处理函数使用
		c.Set("username", claims.Username)
		c.Set("userid", claims.UserId)

		// 将 gin.Context 设置到全局变量中，以便在 service 中使用
		global.GinContext = c

		c.Next()

	}

}
