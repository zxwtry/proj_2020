package comm_constant

const (
	// 函数类
	FUNCTION_PARAM_ERROR          = 001001 // 函数参数出错
	FUNCTION_EXEC_ERROR           = 001002 // 函数执行出错
	FUNCTION_JSON_UNMARSHAL_ERROR = 001003 // json反序列化失败
	FUNCTION_JSON_MARSHAL_ERROR   = 001004 // json序列化失败

	// http simple notoken的请求
	HTTP_SIMPLE_NOTOKEN_NEW_REQUEST_ERROR   = 002001 // 创建request请求失败
	HTTP_SIMPLE_NOTOKEN_DO_REQUEST_ERROR    = 002002 // client.Do失败
	HTTP_SIMPLE_NOTOKEN_READ_RESPONSE_ERROR = 002003 // 从远端读取信息失败

	// url下载类
	URL_DOWNLOAD_REQUEST_ERROR  = 101001 // 通过url下载，请求出错
	URL_DOWNLOAD_DOWNLOAD_ERROR = 101002 // 通过url下载，文件下载出错
	URL_DOWNLOAD_DECODE_ERROR   = 102003 // 通过url下载，解码出错
)
