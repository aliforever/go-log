package log

import "fmt"

func (l *Logger) Alert(a ...interface{}) {
	if l.level == 0 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

func (l *Logger) AlertF(format string, a ...interface{}) {
	if l.level == 0 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

func (l *Logger) Error(a ...interface{}) {
	if l.level == 0 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

func (l *Logger) ErrorF(format string, a ...interface{}) {
	if l.level == 0 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

func (l *Logger) Warn(a ...interface{}) {
	if l.level < 2 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

func (l *Logger) WarnF(format string, a ...interface{}) {
	if l.level < 2 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

func (l *Logger) Highlight(a ...interface{}) {
	if l.level < 3 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

func (l *Logger) HighlightF(format string, a ...interface{}) {
	if l.level < 3 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

func (l *Logger) Inform(a ...interface{}) {
	if l.level < 4 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

func (l *Logger) InformF(format string, a ...interface{}) {
	if l.level < 4 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

func (l *Logger) Log(a ...interface{}) {
	if l.level < 5 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

func (l *Logger) LogF(format string, a ...interface{}) {
	if l.level < 5 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}

func (l *Logger) Trace(a ...interface{}) {
	if l.level < 6 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, a...)...)
}

func (l *Logger) TraceF(format string, a ...interface{}) {
	if l.level < 6 {
		return
	}
	fmt.Fprintln(l.output, append([]interface{}{l.prefix()}, fmt.Sprintf(format, a...))...)
}
