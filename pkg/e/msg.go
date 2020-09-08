package e


var MsgFlags = map[int]string {
	SUCCESS : "请求成功",
	ERROR : "请求失败",
	INVALID_PARAMS : "请求参数错误",


	ERROR_CRON_SPEC : "crontab语法错误",

}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
