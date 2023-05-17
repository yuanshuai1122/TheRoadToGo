package main

import "github.com/gin-gonic/gin"

func main() {

	// 设置cookie
	// func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	//name	string	cookie名字
	//value	string	cookie值
	//maxAge	int	有效时间，单位是秒，MaxAge=0 忽略MaxAge属性，MaxAge<0 相当于删除cookie, 通常可以设置-1代表删除，MaxAge>0 多少秒后cookie失效
	//path	string	cookie路径
	//domain	string	cookie作用域
	//secure	bool	Secure=true，那么这个cookie只能用https协议发送给服务器
	//httpOnly	bool	设置HttpOnly=true的cookie不能被js获取到
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("site_cookie", "cookievalue", 3600, "/", "localhost", false, true)
	})
	// 读取cookie
	r.GET("/read", func(c *gin.Context) {
		// 根据cookie名字读取cookie值
		data, err := c.Cookie("site_cookie")
		if err != nil {
			// 直接返回cookie值
			c.String(200, data)
			return
		}
		c.String(200, "not found!")
	})
	// 删除cookie
	r.GET("/del", func(c *gin.Context) {
		// 设置cookie  MaxAge设置为-1，表示删除cookie
		c.SetCookie("site_cookie", "cookievalue", -1, "/", "localhost", false, true)
		c.String(200, "删除cookie")
	})

	// Session
	// 在Gin框架中，我们可以依赖gin-contrib/sessions中间件处理session。
	r.Run()
}
