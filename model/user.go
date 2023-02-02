package model

import (
	"gin-blog/errormsg"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=2,max=10" label:"用户名" `
	Password string `gorm:"type:varchar(500);not null" json:"password"  validate:"required,min=8,max=18" label:"密码" `
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"权限"`
}

// 获取用户权限
func GetRole(id int) int {
	var user User
	db.Select("*").Where("ID = ?", id).Find(&user)
	if user.ID > 0 {
		return user.Role
	} else {
		return 3
	}
}

// 添加用户
func CreateUser(data *User) int {
	// 加密
	data.Password = scryptPwd(data.Password)
	err := db.Create(data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS

}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS

}

// 编辑用户
func EditUser(id int, data *User) int {
	var user User
	updateMap := make(map[string]interface{}, 4)
	if data.Username != " " {
		updateMap["username"] = data.Username
	}
	updateMap["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(updateMap).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 查询用户列表
func GetUserList(param string, pageSize int, pageNum int) ([]*User, int64) {
	// 用来存储列表的切片
	var users []*User
	// 总数
	var total int64
	// 从数据库中获取
	// 判断-1的情况
	var offset int
	if pageNum == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	err := db.Select("id,username,role").Where("username like ?", param+"%").Limit(pageSize).Offset(offset).Find(&users).Error
	db.Model(&User{}).Where("username like ?", param+"%").Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total

}

// 搜索用户

// 管理员登录
func UserLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).Find(&user)
	// 用户不存在
	if user.ID == 0 {
		return errormsg.ERROR_USER_NOT_EXIST
	}
	if user.Role > 1 {
		return errormsg.ERROR_USER_NO_RIGHT
	}
	PasswordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// 密码错误
	if PasswordErr != nil {
		return errormsg.ERROR_PASSWORD_WRONG
	}
	return errormsg.SUCCESS
}

// 对密码进行加密
func scryptPwd(password string) string {
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
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
