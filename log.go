package log

import (
	"io"
	"os"
	"time"

	"github.com/go-kit/kit/log"
)

var logger = newLog(os.Stderr)

type Log struct{ ctx *log.Context }

func newLog(w io.Writer) *Log {
	return &Log{ctx: log.NewContext(log.NewJSONLogger(w))}
}

func (l *Log) KV(k string, v interface{}) *Log {
	switch s := v.(type) {
	case interface {
		String() string
	}:
		v = s.String()
	case interface {
		GoString() string
	}:
		v = s.GoString()
	}
	return &Log{ctx: l.ctx.With(k, v)}
}

func (l *Log) Err(err error) *Log { return l.KV("err", err) }
func (l *Log) Error(msg string)   { l.log("error", msg) }
func (l *Log) Info(msg string)    { l.log("info", msg) }
func (l *Log) Fatal(msg string)   { l.log("fatal", msg); os.Exit(1) }

func (l *Log) log(lvl, msg string) {
	err := l.ctx.Log(
		"level", lvl,
		"msg", msg,
		"src", log.DefaultCaller(),
		"time", time.Now().UTC().Format(time.RFC3339Nano),
	)
	if err != nil {
		panic(err)
	}
}

func Err(err error) *Log              { return logger.Err(err) }
func Error(msg string)                { logger.log("error", msg) }
func Info(msg string)                 { logger.log("info", msg) }
func Fatal(msg string)                { logger.log("fatal", msg); os.Exit(1) }
func KV(k string, v interface{}) *Log { return logger.KV(k, v) }
