package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000... 用户模块错误
	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WRONG     = 1002
	ERROR_USER_NOT_EXIST     = 1003
	ERROR_TOKEN_NOT_EXIST    = 1004
	ERROE_TOKEN_OUT_OF_TIME  = 1005
	ERROR_TOKEN_WRONG        = 1006
	ERROE_TOKEN_FORMAT_WRONG = 1007
	// code = 2000... 文章模块错误
	// code = 3000... 分类模块错误
)

var codeMsg = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ERROR_USERNAME_USED:      "用户名已存在",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USER_NOT_EXIST:     "用户名不存在",
	ERROR_TOKEN_NOT_EXIST:    "TOKEN不存在",
	ERROE_TOKEN_OUT_OF_TIME:  "TOKEN过期",
	ERROR_TOKEN_WRONG:        "TOKEN不正确",
	ERROE_TOKEN_FORMAT_WRONG: "TOKEN格式错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
