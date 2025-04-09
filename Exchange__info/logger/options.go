package logger

import "log/slog"

type Options func(*PrettyHandlerOptions)

func WithLevel(level slog.Level) Options {
	return func(options *PrettyHandlerOptions) {
		options.SlogOpts.Level = level
	}
}

// 设置时间格式
func WithTimeFormat(format string) Options {
	return func(options *PrettyHandlerOptions) {
		options.TimeFormat = format
	}
}

func WithUseColor(useColor bool) Options {
	return func(options *PrettyHandlerOptions) {
		options.UseColor = useColor
	}
}

func WithOutputJson(outputJson bool) Options {
	return func(options *PrettyHandlerOptions) {
		options.OutPutJson = outputJson
	}
}
