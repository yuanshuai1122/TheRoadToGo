package dao

import "log"

type User struct {
	ID         int64
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
}

func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}

// 插入user
func Save(user *User) {
	err := DB.Create(user)
	if err != nil {
		log.Println("insert fail : ", err)
	}
}

// 根据id获取user
func GetById(id int64) User {
	var user User
	err := DB.Where("id=?", id).First(&user).Error
	if err != nil {
		log.Println("get user by id fail : ", err)
	}
	return user
}

// 获取所有users
func GetAll() []User {
	var users []User
	err := DB.Find(&users)
	if err != nil {
		log.Println("get users  fail : ", err)
	}
	return users
}

// 通过id更新user
func UpdateById(id int64) {
	err := DB.Model(&User{}).Where("id=?", id).Update("username", "lisi")
	if err != nil {
		log.Println("update users  fail : ", err)
	}
}

// 根据id删除user
func DeleteById(id int64) {
	err := DB.Where("id=?", id).Delete(&User{})
	if err != nil {
		log.Println("delete users  fail : ", err)
	}
}
