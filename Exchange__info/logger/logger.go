package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"log/slog"
	"os"
)

var Log *Extend

// 日志配置
type PrettyHandlerOptions struct {
	SlogOpts   slog.HandlerOptions
	TimeFormat string
	UseColor   bool
	OutPutJson bool
}

// 具体的调用
type PrettyHandler struct {
	slog.Handler
	l   *log.Logger
	opt PrettyHandlerOptions
}

// 定义一个总的接口
type Extend struct {
	*slog.Logger
	handler *PrettyHandler
}

func (l *Extend) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

func (l *Extend) DebugMsaf(format string, args ...interface{}) {
	sprintf := fmt.Sprintf(format, args...)
	l.Logger.Debug(sprintf)
}

func (l *Extend) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *Extend) InfoMsaf(format string, args ...interface{}) {
	sprintf := fmt.Sprintf(format, args...)
	l.Logger.Info(sprintf)
}

func (l *Extend) Warn(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}

func (l *Extend) WarnMsaf(format string, args ...interface{}) {
	sprintf := fmt.Sprintf(format, args...)
	l.Logger.Warn(sprintf)
}

func (l *Extend) Error(msg string, args ...any) error {
	l.Logger.Error(msg, args...)
	return fmt.Errorf(msg)
}

func (l *Extend) ErrorMsaf(format string, args ...interface{}) error {
	sprintf := fmt.Sprintf(format, args...)
	l.Logger.Error(sprintf)
	return fmt.Errorf(sprintf)
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	//从这边获取 后面实现参数
	TimeStr := r.Time.Format(h.opt.TimeFormat)
	level := "[" + r.Level.String() + "]:"
	//判断颜色类型
	if h.opt.UseColor {
		switch r.Level {
		case slog.LevelDebug:
			level = color.MagentaString(level)
		case slog.LevelInfo:
			level = color.BlueString(level)
		case slog.LevelWarn:
			level = color.YellowString(level)
		case slog.LevelError:
			level = color.RedString(level)
		}
	}
	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(attr slog.Attr) bool {
		fields[attr.Key] = attr.Value.Any()
		return true
	})
	//判断输出是否是json
	if len(fields) > 0 && h.opt.OutPutJson {
		indent, err := json.MarshalIndent(fields, "", "")
		if err != nil {
			return err
		}
		h.l.Printf("%s %s %s %s\n", TimeStr, level, r.Message, string(indent))
	} else {
		h.l.Printf("%s %s %s\n", TimeStr, level, r.Message)
	}
	return nil
}

func NewPrettyHandler(out io.Writer, opts PrettyHandlerOptions) *PrettyHandler {
	p := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
		opt:     opts,
	}
	return p
}

func Init(option ...Options) {
	options := PrettyHandlerOptions{
		SlogOpts:   slog.HandlerOptions{Level: slog.LevelDebug},
		TimeFormat: "2006-01-02 15:04:05",
		UseColor:   true,
		OutPutJson: true,
	}
	//上面是默认的 如果想修改 需要循环
	for _, opts := range option {
		opts(&options)
	}
	handler := NewPrettyHandler(os.Stdout, options)
	Log = &Extend{
		Logger:  slog.New(handler),
		handler: handler,
	}

}
