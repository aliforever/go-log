package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	output   io.Writer
	prefixes []string
	beganAt  int64
	level    uint8
}

func NewLogger(writer io.Writer) *Logger {
	if writer == nil {
		writer = os.Stdout
	}
	return &Logger{output: writer}
}

func (l *Logger) Begin() *Logger {
	fmt.Fprintln(l.output, "BEGIN")
	l.beganAt = time.Now().UnixNano()
	return &Logger{
		beganAt:  time.Now().UnixNano(),
		output:   l.output,
		prefixes: l.prefixes,
	}
}

func (l *Logger) End() {
	fmt.Fprintln(l.output, "END", fmt.Sprintf("&t=%dµs", int64(float64(time.Now().UnixNano()-l.beganAt)*0.001)))
}

func (l *Logger) trace() string {
	pc, _, _, ok := runtime.Caller(3)
	if !ok {
		return ""
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "?"
	}

	return fn.Name()
}

func (l *Logger) prefix() interface{} {
	if len(l.prefixes) > 0 {
		return l.trace() + ": " + strings.Join(l.prefixes, ": ") + ":"
	}
	return l.trace() + ":"
}

func (l *Logger) Prefix(prefixes ...string) *Logger {
	return &Logger{
		output:   l.output,
		prefixes: append(l.prefixes, prefixes...),
	}
}

func (l *Logger) Level(level uint8) *Logger {
	return &Logger{
		output:   l.output,
		prefixes: l.prefixes,
		beganAt:  time.Now().UnixNano(),
		level:    level,
	}
}
