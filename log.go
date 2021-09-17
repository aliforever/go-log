package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	output   io.Writer
	prefixes []string
}

func NewLogger(writer io.Writer) *Logger {
	if writer == nil {
		writer = os.Stdout
	}
	return &Logger{output: writer}
}

func (l *Logger) Begin() {
	fmt.Fprintln(l.output, "BEGIN")
}

func (l *Logger) End() {
	fmt.Fprintln(l.output, "END")
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

func (l *Logger) Log(a ...interface{}) {
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

func (l *Logger) LogF(format string, a ...interface{}) {
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

func (l *Logger) Prefix(prefixes ...string) *Logger {
	return &Logger{
		output:   l.output,
		prefixes: append(l.prefixes, prefixes...),
	}
}
