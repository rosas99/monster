package known

const (
	// TraceIDKey 用来定义 Gin 上下文中的键，代表请求的 uuid.
	TraceIDKey = "Trace-ID"

	// UsernameKey 用来定义 Gin 上下文的键，代表请求的所有者.
	UsernameKey = "Username"
)
