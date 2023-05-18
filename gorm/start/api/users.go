package api

import (
	"TheRoadToGo/gorm/start/dao"
	"github.com/gin-gonic/gin"
	"time"
)

func SaveUser(c *gin.Context) {
	user := &dao.User{
		Username:   "zhangsan",
		Password:   "123456",
		CreateTime: time.Now().UnixMilli(),
	}
	dao.Save(user)
	c.JSON(200, user)
}

func RegisterRouter(r *gin.Engine) {
	r.GET("/save", SaveUser)
}

// 。。。。

func GetUser(c *gin.Context) {
	user := dao.GetById(1)
	c.JSON(200, user)
}

func GetUserAll(c *gin.Context) {
	user := dao.GetAll()
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	dao.UpdateById(1)
	user := dao.GetById(1)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	dao.DeleteById(1)
	user := dao.GetById(1)
	c.JSON(200, user)
}
