package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/go-programming-tour-book/blog-service/pkg/limiter"
)

//  限流控制
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {

		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			// 占用存储桶中立即可用的令牌的数量，返回值为删除的令牌数，如果没有可用的令牌，将会返回 0，也就是已经超出配额了
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				// 将返回 errcode.TooManyRequest 状态告诉客户端需 要减缓并控制请求速度
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
