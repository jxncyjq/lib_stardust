//公共错误库
package errors

// 0-199	预留
// 200-299  参数错误
// 300-399	类型错误
// 400-499  中间件问题
// 500-599  内部错误
// 600-699  操作系统错误
// 700-699  协议错误

var (
	// ErrUnsupportedContentType indicates unacceptable or lack of Content-Type
	ErrUnsupportedContentType = New("unsupported content type",300)

	// ErrInvalidQueryParams indicates invalid query parameters
	ErrInvalidQueryParams = New("invalid query parameters",200)

	// ErrNotFoundParam indicates that the parameter was not found in the query
	ErrNotFoundParam = New("parameter not found in the query",201)

	// ErrMalformedEntity indicates a malformed entity specification
	ErrMalformedEntity = New("malformed entity specification",500)


)

var (
	ErrRedisConnection = New("redis连接错误",400)
	ErrRedisConfig = New("Redis配置错误",401)
	ErrDatabaseConnect = New("数据库连接错误",403)
	ErrDatabaseSessionNil = New("数据库会话为空",405)
	ErrDatabaseCommit = New("数据库提交错误",406)
)

var (
	ErrAuthConnection = New("Auth Grpc 连接错误",404)
	ErrCloseTracer = New("关闭Tracer错误",407)
)
