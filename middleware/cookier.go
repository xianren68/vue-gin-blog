package middleware

import (
	"gin-blog/errormsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cookier() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取cookie
		role, err := c.Cookie("role")

		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"status": errormsg.ERROR_COOKIE_RUNTIME,
				"msg":    errormsg.GetMsg(errormsg.ERROR_COOKIE_RUNTIME),
			})
		} else {
			c.Set("role", role)
		}
	}

}
