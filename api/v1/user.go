package v1

import (
	"gin-blog/errormsg"
	"gin-blog/model"
	"gin-blog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

// 定义一个结构体
type UserApi struct {
}

// 添加用户
func (u *UserApi) Add(c *gin.Context) {
	var data model.User
	// 将用户数据绑定到结构体
	c.ShouldBindJSON(&data)
	// 进行验证
	msg, code := validator.Validate(&data)
	if code == errormsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    msg,
		})
		c.Abort()
		return
	}
	code = model.UserExist(data.Username)
	// 如果不存在，则添加到数据库
	if code == errormsg.SUCCESS {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
	})
}

// 删除用户
func (u *UserApi) Delete(c *gin.Context) {
	// 获取id
	id, _ := strconv.Atoi(c.Param("id"))
	// 数据库操作
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
	})

}

// 编辑用户
func (u *UserApi) Edit(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&user)
	// 判断要修改的用户名是否已存在
	code = model.UserExist(user.Username)
	if code == errormsg.SUCCESS {
		code = model.EditUser(id, &user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
	})
}

// 查询用户列表
func (u *UserApi) List(c *gin.Context) {
	// 转换为int类型
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	// 如果没有传参设置为-1，返回所有数据
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	// 从数据库中获取
	data, total := model.GetUserList(pageSize, pageNum)
	code = errormsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errormsg.GetMsg(code),
		"total":  total,
	})

}
