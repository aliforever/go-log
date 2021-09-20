package tests

import (
	"testing"

	"github.com/aliforever/go-log"
)

func TestLogger_Log_Level_Zero(t *testing.T) {
	l := log.NewLogger(nil).Level(0)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")
}

func TestLogger_Log_Level_One(t *testing.T) {
	l := log.NewLogger(nil).Level(1)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")
}

func TestLogger_Log_Level_Two(t *testing.T) {
	l := log.NewLogger(nil).Level(2)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")
}

func TestLogger_Log_Level_Three(t *testing.T) {
	l := log.NewLogger(nil).Level(3)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")
}

func TestLogger_Log_Level_Four(t *testing.T) {
	l := log.NewLogger(nil).Level(4)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")
}

func TestLogger_Log_Level_Five(t *testing.T) {
	l := log.NewLogger(nil).Level(5)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")
}

func TestLogger_Log_Level_Six(t *testing.T) {
	l := log.NewLogger(nil).Level(6)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")
}

func TestLogger_Log_Level_Upper(t *testing.T) {
	l := log.NewLogger(nil).Level(7)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")
}
