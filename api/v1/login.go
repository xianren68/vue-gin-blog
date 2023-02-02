package v1

import (
	"gin-blog/errormsg"
	"gin-blog/middleware"
	"gin-blog/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginApi struct {
}

// 后台登录
func (L *LoginApi) UserLogin(c *gin.Context) {
	var user model.User
	// 绑定到结构体
	c.ShouldBind(&user)
	code = model.UserLogin(user.Username, user.Password)
	token := ""
	if code == errormsg.SUCCESS {
		// 设置token
		token, code = middleware.CreateToken(user.Username)
		// 设置cookie,携带其role
		c.SetCookie("role", strconv.Itoa(user.Role), 10*60*60, "/", "localhost", false, false)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
		"token":  token,
	})
}
