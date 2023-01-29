package middleware

import (
	"gin-blog/errormsg"
	"gin-blog/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte(utils.JwtKye)
var code int

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 创建token
func CreateToken(username string) (string, int) {
	// 设置过期时间
	expireTime := time.Now().Add(10 * time.Hour)
	setClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", errormsg.ERROR
	}
	return token, errormsg.SUCCESS

}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	// 解析
	settoken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, code := settoken.Claims.(*MyClaims); code && settoken.Valid {
		return key, errormsg.SUCCESS
	} else {
		return nil, errormsg.ERROR
	}
}

// 编写token中间件

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code = errormsg.SUCCESS
		// 不存在token
		if tokenHeader == "" {
			code = errormsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errormsg.GetMsg(code),
			})
			c.Abort()
			return
		}
		checktoken := strings.SplitN(tokenHeader, " ", 2)
		// 格式出错
		if len(checktoken) != 2 || checktoken[0] != "Bearer" {
			code = errormsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errormsg.GetMsg(code),
			})
			c.Abort()
			return
		}
		key, flag := CheckToken(checktoken[1])
		// 解析出错
		if flag == errormsg.ERROR {
			code = errormsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errormsg.GetMsg(code),
			})
			c.Abort()
			return
		}
		// 超时
		if key.ExpiresAt < time.Now().Unix() {
			code = errormsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errormsg.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
	}
}
