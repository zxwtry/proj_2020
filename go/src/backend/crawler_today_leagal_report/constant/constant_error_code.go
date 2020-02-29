package constant

const (
	// 函数类
	FUNCTION_PARAM_ERROR = 001001 //函数参数出错

	// http simple notoken的请求
	HTTP_SIMPLLE_NOTOKEN_NEW_REQUEST = 002001 // 创建request请求失败

	// url下载类
	URL_DOWNLOAD_REQUEST_ERROR  = 101001 // 通过url下载，请求出错
	URL_DOWNLOAD_DOWNLOAD_ERROR = 101002 // 通过url下载，文件下载出错
	URL_DOWNLOAD_DECODE_ERROR   = 102003 // 通过url下载，解码出错
)
