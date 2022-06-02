package xerr

// 通用状态码
const (
	OK = 0 // 成功

	CLIENT_ERROR = 4000 // 客户端自身错误
	UNAUTHORIZED = 4001 // 身份认证失败
	PARAM_ERROR  = 4002 // 缺少必传参数
	NO_DATA      = 4003 // 没有数据
	DATA_EXIST   = 4004 // 数据已存在

	SERVER_ERROR  = 5000 // 服务端开小差啦,请稍后再试一试
	DB_ERROR      = 5001 // 数据库操作异常
	CACHE_ERROR   = 5002 // 缓存操作异常
	MARSHAL_ERROR = 5003 // 序列化异常
	MQ_ERROR      = 5004 // 消息队列操作异常
	WS_ERROR	= 5005 // websocket操作异常
)

