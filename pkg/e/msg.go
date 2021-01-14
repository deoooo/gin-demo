package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "server error",
	INVALID_PARAMS: "错误的参数",
	NEED_AUTH:      "请登陆",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
