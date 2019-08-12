package e

const (
	SUCCESS = 200 // 成功
	ERROR = 500 // 服务器错误
	INVALID_PARAMS = 400 // 非法参数

	ERROR_EXIST_TAG = 10001 // 错误标签
	ERROR_NOT_EXIST_TAG = 10002 // 未存在标签
	ERROR_NOT_EXIST_ARTICLE = 10003 // 未存在文章

	ERROR_AUTH_CHECK_TOKEN_FAIL = 20001 // 鉴权失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002 // 鉴权超时
	ERROR_AUTH_TOKEN = 20003
	ERROR_AUTH = 20004
)
