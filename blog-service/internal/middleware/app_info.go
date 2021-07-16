package middleware

import "github.com/gin-gonic/gin"

// 需要在进程内上下文设置一些内部信息，
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog-service")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
