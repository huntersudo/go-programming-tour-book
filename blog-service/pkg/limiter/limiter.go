package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	ratelimit "github.com/juju/ratelimit"
)
// 限流器是存在多种实现的，可能某一类接口需要限流器 A， 另外一类接口需要限流器 B，所采用的策略不是完全一致的
type LimiterIface interface {
	Key(c *gin.Context) string // 获取对应的限流器的键值对名称。
	GetBucket(key string) (*ratelimit.Bucket, bool) 	// 获取令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface // 新增多个令牌桶
}
// 用于存储令牌桶与键值对名称的映射关系
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}
 // 结构体用于存储令牌桶的一些相应规则属性
type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration  // 间隔多久时间放 N 个令牌
	Capacity     int64 // 令牌桶的容量
	Quantum      int64  // 每次到达间隔时间后所放的具体令牌数量
}
