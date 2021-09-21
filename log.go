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

// This is the constructor for the Logger
// Writer can be any writers (file, builders, etc...) and defaults to os.Stdout
func NewLogger(writer io.Writer) *Logger {
	if writer == nil {
		writer = os.Stdout
	}
	return &Logger{output: writer}
}

// Begin sets the start time of the logger so later you can calculate the time between BEGIN & END of the log
func (l *Logger) Begin() *Logger {
	fmt.Fprintln(l.output, "BEGIN")
	l.beganAt = time.Now()
	return &Logger{
		beganAt:  time.Now(),
		output:   l.output,
		prefixes: l.prefixes,
		level:    l.level,
	}
}

// End sets the finish time of the logger so later you can calculate the time between BEGIN & END of the log
func (l *Logger) End() {
	fmt.Fprintln(l.output, "END", fmt.Sprintf("&t=%dµs", time.Now().Sub(l.beganAt).Microseconds()))
}

// trace is used to return current main function name
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

// prefix returns the prefixes for the logs
func (l *Logger) prefix() interface{} {
	if len(l.prefixes) > 0 {
		return l.trace() + ": " + strings.Join(l.prefixes, ": ") + ":"
	}
	return l.trace() + ":"
}

// Returns a new Logger with custom prefixes passed
func (l *Logger) Prefix(prefixes ...string) *Logger {
	return &Logger{
		output:   l.output,
		prefixes: append(l.prefixes, prefixes...),
	}
}

// Returns a logger with custom Level (level >= 0)
// Level0: None
// Level1: Alert & Error
// Level2: Level1 & Warn
// Level3: Level2 & Highlight
// Level4: Level3 & Inform
// Level5: Level4 & Log
// Level6: Level5 & Trace
func (l *Logger) Level(level uint8) *Logger {
	return &Logger{
		output:   l.output,
		prefixes: l.prefixes,
		beganAt:  time.Now(),
		level:    level,
	}
}
