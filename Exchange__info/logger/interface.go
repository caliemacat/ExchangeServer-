package logger

// 新建接口调用日志
type Ilogger interface {
	Debug(msg string, args ...any)
	DebugMsaf(format string, args ...interface{})

	Info(msg string, args ...any)
	InfoMsaf(format string, args ...interface{})
	Warn(msg string, args ...any)
	WarnMsaf(format string, args ...interface{})
	Error(msg string, args ...any)
	ErrorMsaf(format string, args ...interface{})
}
