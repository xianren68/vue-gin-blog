package model

import (
	"gin-blog/errormsg"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" `
	Password string `gorm:"type:varchar(500);not null" json:"password" `
	Role     int    `gorm:"type:int" json:"role" `
}

// 判断用户名是否存在
func UserExist(username string) int {
	var user User
	// 查询用户名是否存在于数据库中,并将值返回到user中
	db.Select("ID").Where("username = ?", username).Find(&user)

	// 用户已存在
	if user.ID > 0 {
		return errormsg.ERROR_USER_USED
	}
	return errormsg.SUCCESS
}

// 添加用户到数据库中
func CreateUser(data *User) int {
	err := db.Create(data)
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS

}

// 查询用户列表
func GetUserList(pageSize int, pageNum int) []*User {
	// 用来存储列表的切片
	var users []*User
	// 从数据库中获取
	// 判断-1的情况
	var offset int
	if pageNum == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	err := db.Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users

}
