package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 我们调用了 context.WithTimeout 方法设置当前 context 的超时时间，
		// 并重新 赋予给了 gin.Context，这样子在当前请求运行到指定的时间后，
		// 在使用了该 context 的运行流 程就会针对 context 所提供的超时时间进行处理
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
