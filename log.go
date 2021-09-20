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
	beganAt  time.Time
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
	l.beganAt = time.Now()
	return &Logger{
		beganAt:  time.Now(),
		output:   l.output,
		prefixes: l.prefixes,
	}
}

func (l *Logger) End() {
	fmt.Fprintln(l.output, "END", fmt.Sprintf("&t=%dµs", time.Now().Sub(l.beganAt).Microseconds()))
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
		beganAt:  time.Now(),
		level:    level,
	}
}
