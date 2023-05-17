package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	// 普通参数
	//http://localhost:8080/user/save?id=11&name=zhangsan
	r.GET("/user/save", func(ctx *gin.Context) {
		id := ctx.Query("id")
		name := ctx.Query("name")
		ctx.JSON(200, gin.H{
			"id":   id,
			"name": name,
		})
	})
	// 如果参数不存在，就给一个默认值
	r.GET("/user/save1", func(ctx *gin.Context) {
		id := ctx.Query("id")
		name := ctx.Query("name")
		address := ctx.DefaultQuery("address", "北京")
		ctx.JSON(200, gin.H{
			"id":      id,
			"name":    name,
			"address": address,
		})
	})
	// 判断参数是否存在
	// localhost:8080/user/save2
	// {"address":"","aok":false,"id":"","idok":false}
	r.GET("/user/save2", func(ctx *gin.Context) {
		id, ok := ctx.GetQuery("id")
		address, aok := ctx.GetQuery("address")
		ctx.JSON(200, gin.H{
			"id":      id,
			"idok":    ok,
			"address": address,
			"aok":     aok,
		})
	})

	// id是数值类型，上述获取的都是string类型，根据类型获取
	type User struct {
		Id   int64  `form:"id"`
		Name string `form:"name"`
	}
	r.GET("/user/save3", func(ctx *gin.Context) {
		var user User
		err := ctx.BindQuery(&user)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(200, user)
	})
	// 也可以：
	r.GET("/user/save4", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBindQuery(&user)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(200, user)
	})
	// 区别
	// type User struct {
	//	Id      int64  `form:"id"`
	//	Name    string `form:"name"`
	//	Address string `form:"address" binding:"required"`
	//}
	// 当bind是必须的时候，ShouldBindQuery会报错，开发者自行处理，状态码不变。
	//BindQuery则报错的同时，会将状态码改为400。所以一般建议是使用Should开头的bind。

	r.Run()

}
