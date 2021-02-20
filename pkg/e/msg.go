package e


var MsgFlags = map[int]string {
	SUCCESS : "请求成功",
	ERROR : "请求失败",
	INVALID_PARAMS : "请求参数错误",
	INVALID_PASS : "旧密码错误",
	INVALID_DIFFPASS : "两次新密码不一致",
	INVALID_FINGER : "需要去掉指纹尾部逗号",


	ERROR_CRON_SPEC : "crontab语法错误",

}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
