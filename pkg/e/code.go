package e

const (
	SUCCESS                         = 0
	ERROR                           = 500
	INVALID_PARAMS                  = 1001
	ERROR_AUTH_CHECK_TOKEN_FAIL     = 4001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT  = 4002
	ERROR_AUTH_TOKEN                = 4003
	ERROR_AUTH_PARAMS               = 4004
	ERROR_AUTH_TOKEN_EMPTY          = 4005
	ERROR_USER_NOT_EXIST            = 4006
	ERROR_REGISTER_FAILED           = 4007
	ERROR_USERNAME_EXIST            = 4008
	ERROR_NICKNAME_EXIST            = 4009
	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 2001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 2002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 2003
)

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "登录状态无效",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "登录状态已超时",
	ERROR_AUTH_TOKEN:                "登录状态生成失败",
	ERROR_AUTH_PARAMS:               "登录失败,账号或密码不正确",
	ERROR_AUTH_TOKEN_EMPTY:          "未登录",
	ERROR_USER_NOT_EXIST:            "用户不存在",
	ERROR_REGISTER_FAILED:           "注册失败",
	ERROR_USERNAME_EXIST:            "用户名已存在",
	ERROR_NICKNAME_EXIST:            "昵称已存在",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
