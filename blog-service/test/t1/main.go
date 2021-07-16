package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()
}
// 默认 Engine 实例：当前默认使用了官方所提供的 Logger 和 Recovery 中间件创建了 Engine 实例。
// 运行模式：当前为调试模式，并建议若在生产环境时切换为发布模式。
// 路由注册：注册了 GET /ping 的路由，并输出其调用方法的方法名。
// 运行信息：本次启动时监听 8080 端口，由于没有设置端口号等信息，因此默认为 8080


// $ go run main.go
// [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
//
//[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
// - using env:   export GIN_MODE=release
// - using code:  gin.SetMode(gin.ReleaseMode)

//[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
//...
//[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
//[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
//[GIN-debug] Listening and serving HTTP on :8080