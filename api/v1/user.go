package v1

import (
	"fmt"
	"gin-blog/errormsg"
	"gin-blog/model"
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

// 修改密码
func (u *UserApi) ChangePass(c *gin.Context) {

}

// 查询单个用户
func (u *UserApi) GetOnly(c *gin.Context) {

}

// 查询用户列表
func (u *UserApi) GetList(c *gin.Context) {
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
	fmt.Println(pageNum)
	// 从数据库中获取
	data := model.GetUserList(pageSize, pageNum)
	code = errormsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errormsg.GetMsg(code),
	})

}
