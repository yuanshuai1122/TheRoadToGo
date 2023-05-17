package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 请求方法
	r.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(200, "get")
	})
	r.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(200, "post")
	})
	r.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(200, "delete")
	})
	r.PUT("/put", func(ctx *gin.Context) {
		ctx.JSON(200, "put")
	})

	//如果想要支持所有
	r.Any("/any", func(ctx *gin.Context) {
		ctx.JSON(200, "any")
	})

	// 如果想要支持其中的几种
	r.GET("/hello", func(ctx *gin.Context) {
		//数组 map list 结构体
		ctx.JSON(200, gin.H{
			"name": "hello world",
		})
	})
	r.POST("/hello", func(ctx *gin.Context) {
		//数组 map list 结构体
		ctx.JSON(200, gin.H{
			"name": "hello world",
		})
	})

	// URI
	//静态url，比如/hello，/user/find
	r.POST("/user/find", func(ctx *gin.Context) {
	})
	//路径参数，比如/user/find/:id
	r.POST("/user/find/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		ctx.JSON(200, param)
	})
	// 模糊匹配，比如/user/*path
	//r.POST("/user/*path", func(ctx *gin.Context) {
	//	param := ctx.Param("path")
	//	ctx.JSON(200, param)
	//})
	// 处理函数
	// type HandlerFunc func(*Context)
	// 通过上下文的参数，获取http的请求参数，响应http请求等。

	// 分组路由
	// 在进行开发的时候，我们往往要进行模块的划分，比如用户模块，以user开发，商品模块，以goods开头。
	//或者进行多版本开发，不同版本之间路径是一致的，这种时候，就可以用到分组路由了。
	ug := r.Group("/user")
	{
		ug.GET("find", func(ctx *gin.Context) {
			ctx.JSON(200, "user find")
		})
		ug.POST("save", func(ctx *gin.Context) {
			ctx.JSON(200, "user save")
		})
	}
	gg := r.Group("/goods")
	{
		gg.GET("find", func(ctx *gin.Context) {
			ctx.JSON(200, "goods find")
		})
		gg.POST("save", func(ctx *gin.Context) {
			ctx.JSON(200, "goods save")
		})
	}
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
