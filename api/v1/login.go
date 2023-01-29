package v1

import (
	"gin-blog/errormsg"
	"gin-blog/middleware"
	"gin-blog/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginApi struct {
}

func (L *LoginApi) UserLogin(c *gin.Context) {
	var user model.User
	// 绑定到结构体
	c.ShouldBindJSON(&user)
	code = model.UserLogin(user.Username, user.Password)
	token := ""
	if code == errormsg.SUCCESS {
		// 设置token
		token, code = middleware.CreateToken(user.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
		"token":  token,
	})
}
