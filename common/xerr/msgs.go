package xerr

var mapCodeMsg = map[uint32]string{
	OK: "成功",

	CLIENT_ERROR: "客户端自身错误",
	UNAUTHORIZED: "身份认证失败",
	PARAM_ERROR:  "缺少必传参数",
	NO_DATA:      "没有数据",
	DATA_EXIST:   "数据已存在",

	SERVER_ERROR: "服务端开小差啦,请稍后再试一试",
	DB_ERROR:     "数据库操作异常",
	CACHE_ERROR:  "缓存操作异常",
	MARSHAL_ERROR: "序列化异常",
	MQ_ERROR: "消息队列操作异常",
	WS_ERROR: "websocket操作异常",
}

func CodeToMsg(code uint32) string {
	if msg, ok := mapCodeMsg[code]; ok {
		return msg
	} else { // 没找到, 则返回默认的错误
		return mapCodeMsg[SERVER_ERROR]
	}
}

func IsCodeErr(code uint32) bool {
	if _, ok := mapCodeMsg[code]; ok {
		return true
	} else {
		return false
	}
}
