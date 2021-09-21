package log

import "fmt"

// These are the levels to the Logger
// If you set the Log Level to 0, none would output

// Alert is used for alerts and writes to Logger's output
// Log Level is 1
func (l *Logger) Alert(a ...interface{}) {
	if l.level == 0 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

// AlertF is used for alerts and receives a format to use and writes to Logger's output
// Log Level is 1
func (l *Logger) AlertF(format string, a ...interface{}) {
	if l.level == 0 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

// Error is used for alerts and writes to Logger's output
// Log Level is 1
func (l *Logger) Error(a ...interface{}) {
	if l.level == 0 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

// ErrorF is used for alerts and receives a format to use and writes to Logger's output
// Log Level is 1
func (l *Logger) ErrorF(format string, a ...interface{}) {
	if l.level == 0 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

// Warn is used for alerts and writes to Logger's output
// Log Level is 2
func (l *Logger) Warn(a ...interface{}) {
	if l.level < 2 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

// WarnF is used for alerts and receives a format to use and writes to Logger's output
// Log Level is 2
func (l *Logger) WarnF(format string, a ...interface{}) {
	if l.level < 2 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

// Highlight is used for alerts and writes to Logger's output
// Log Level is 3
func (l *Logger) Highlight(a ...interface{}) {
	if l.level < 3 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

// HighlightF is used for alerts and receives a format to use and writes to Logger's output
// Log Level is 3
func (l *Logger) HighlightF(format string, a ...interface{}) {
	if l.level < 3 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

// Inform is used for alerts and writes to Logger's output
// Log Level is 4
func (l *Logger) Inform(a ...interface{}) {
	if l.level < 4 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

// InformF is used for alerts and receives a format to use and writes to Logger's output
// Log Level is 4
func (l *Logger) InformF(format string, a ...interface{}) {
	if l.level < 4 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

// Log is used for alerts and writes to Logger's output
// Log Level is 5
func (l *Logger) Log(a ...interface{}) {
	if l.level < 5 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

// LogF is used for alerts and receives a format to use and writes to Logger's output
// Log Level is 5
func (l *Logger) LogF(format string, a ...interface{}) {
	if l.level < 5 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

// Trace is used for alerts and writes to Logger's output
// Log Level is 6
func (l *Logger) Trace(a ...interface{}) {
	if l.level < 6 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

// TraceF is used for alerts and receives a format to use and writes to Logger's output
// Log Level is 6
func (l *Logger) TraceF(format string, a ...interface{}) {
	if l.level < 6 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}
