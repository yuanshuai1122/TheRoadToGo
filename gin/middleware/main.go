package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 自定义个日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 可以通过上下文对象，设置一些依附在上下文对象里面的键/值数据
		c.Set("example", "12345")

		// 在这里处理请求到达控制器函数之前的逻辑

		// 调用下一个中间件，或者控制器处理函数，具体得看注册了多少个中间件。
		c.Next()

		// 在这里可以处理请求返回给用户之前的逻辑
		latency := time.Since(t)
		log.Print(latency)

		// 例如，查询请求状态吗
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {

	// 在Gin框架中，中间件（Middleware）指的是可以拦截http请求-响应生命周期的特殊函数，在请求-响应生命周期中可以注册多个中间件，每个中间件执行不同的功能，一个中间执行完再轮到下一个中间件执行。
	//中间件的常见应用场景如下：
	//请求限速
	//api接口签名处理
	//权限校验
	//统一错误处理
	//Gin支持设置全局中间件和针对路由分组设置中间件，设置全局中间件意思就是会拦截所有请求，针对分组路由设置中间件，意思就是仅对这个分组下的路由起作用。
	r := gin.New()
	// 通过use设置全局中间件
	// 设置日志中间件，主要用于打印请求日志
	r.Use(Logger())
	// 设置Recovery中间件，主要用于拦截paic错误，不至于导致进程崩掉
	r.Use(gin.Recovery())
	r.GET("/test", func(ctx *gin.Context) {
		panic(errors.New("test error"))
	})
	r.Run(":8080")
}
