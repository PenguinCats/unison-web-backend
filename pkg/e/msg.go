package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "内部错误",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:   "登录超时",
	ERROR_AUTH_CHECK_TOKEN_FAIL:      "登录信息无效",
	ERROR_AUTH_PERMISSION_DENIED:     "权限不足",
	ERROR_AUTH_LOGIN_WRONG_UNAME_PWD: "用户名或密码错误",
	ERROR_AUTH_LOGIN_INVALID_SEU_ID:  "该学号尚未绑定有效账号",
	ERROR_AUTH_LOGIN_INVALID_QQ_ID:   "该QQ号尚未绑定有效账号",

	ERROR_OPERATION_GENERATION: "操作指令生成失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
