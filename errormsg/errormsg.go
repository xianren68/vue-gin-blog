package errormsg

// 定义一些常量，用做错误状态码
const (
	SUCCESS = 200
	ERROR   = 500
	// 关于用户数据的错误 1000
	ERROR_USER_USED        = 1001
	ERROR_USER_NOT_EXIST   = 1002
	ERROR_PASSWORD_WRONG   = 1003
	ERROR_USER_NO_RIGHT    = 1004
	ERROR_TOKEN_NOT_EXIST  = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_RUNTIME    = 1007
	ERROR_TOKEN_TYPE_WRONG = 1008
	ERROR_COOKIE_RUNTIME   = 1009
	// 关于文章的错误 2000
	ERROR_ART_NOT_EXIST = 2001
	// 关于分类模块的错误 3000
	ERROR_CATEGORY_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

// 建立错误信息与错误码的映射表
var codeMsg = map[int]string{
	SUCCESS:                "ok",
	ERROR:                  "fail",
	ERROR_USER_USED:        "用户名被占用",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
	ERROR_PASSWORD_WRONG:   "密码有误",
	ERROR_TOKEN_NOT_EXIST:  "token不存在,请重新登录",
	ERROR_TOKEN_WRONG:      "token有误,请重新登录",
	ERROR_TOKEN_RUNTIME:    "token过期,请重新登录",
	ERROR_TOKEN_TYPE_WRONG: "token格式有误,请重新登录",
	ERROR_ART_NOT_EXIST:    "文章不存在",
	ERROR_CATEGORY_USED:    "已有此分类",
	ERROR_CATE_NOT_EXIST:   "没有这个分类",
	ERROR_COOKIE_RUNTIME:   "cookie失效",
}

// 给外面的接口,通过错误码获取msg
func GetMsg(code int) string {
	return codeMsg[code]
}
