package log

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	output io.Writer
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

func (l *Logger) Log(title, message string) {
	fmt.Fprintln(l.output, fmt.Sprintf("%s %s", title, message))
}

func (l *Logger) LogF(format string, a ...interface{}) {
	fmt.Fprintln(l.output, fmt.Sprintf(format, a...))
}
