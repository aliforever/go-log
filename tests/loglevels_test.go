package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aliforever/go-log"
)

func TestLogger_Log_Level_Zero(t *testing.T) {
	var builder strings.Builder

	l := log.NewLogger(&builder).Level(0)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")

	expected := ""
	value := builder.String()

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}
}

func TestLogger_Log_Level_One(t *testing.T) {
	var builder strings.Builder

	l := log.NewLogger(&builder).Level(1)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")

	expected := "github.com/aliforever/go-log/tests.TestLogger_Log_Level_One: This is an alert\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_One: This is an error\n"
	value := builder.String()

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}
}

func TestLogger_Log_Level_Two(t *testing.T) {
	var builder strings.Builder

	l := log.NewLogger(&builder).Level(2)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")

	expected := "github.com/aliforever/go-log/tests.TestLogger_Log_Level_Two: This is an alert\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Two: This is an error\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Two: This is a warning\n"
	value := builder.String()

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}
}

func TestLogger_Log_Level_Three(t *testing.T) {
	var builder strings.Builder

	l := log.NewLogger(&builder).Level(3)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")

	expected := "github.com/aliforever/go-log/tests.TestLogger_Log_Level_Three: This is an alert\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Three: This is an error\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Three: This is a warning\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Three: This is a highlight\n"
	value := builder.String()

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}
}

func TestLogger_Log_Level_Four(t *testing.T) {
	var builder strings.Builder

	l := log.NewLogger(&builder).Level(4)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")

	expected := "github.com/aliforever/go-log/tests.TestLogger_Log_Level_Four: This is an alert\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Four: This is an error\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Four: This is a warning\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Four: This is a highlight\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Four: This is an information\n"
	value := builder.String()

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}
}

func TestLogger_Log_Level_Five(t *testing.T) {
	var builder strings.Builder

	l := log.NewLogger(&builder).Level(5)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")

	expected := "github.com/aliforever/go-log/tests.TestLogger_Log_Level_Five: This is an alert\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Five: This is an error\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Five: This is a warning\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Five: This is a highlight\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Five: This is an information\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Five: This is a log\n"
	value := builder.String()

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}
}

func TestLogger_Log_Level_Six(t *testing.T) {
	var builder strings.Builder

	l := log.NewLogger(&builder).Level(6)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")

	expected := "github.com/aliforever/go-log/tests.TestLogger_Log_Level_Six: This is an alert\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Six: This is an error\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Six: This is a warning\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Six: This is a highlight\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Six: This is an information\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Six: This is a log\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Six: This is a trace\n"
	value := builder.String()

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}
}

func TestLogger_Log_Level_Upper(t *testing.T) {
	var builder strings.Builder

	l := log.NewLogger(&builder).Level(7)
	l.Alert("This is an alert")
	l.Error("This is an error")
	l.Warn("This is a warning")
	l.Highlight("This is a highlight")
	l.Inform("This is an information")
	l.Log("This is a log")
	l.Trace("This is a trace")

	expected := "github.com/aliforever/go-log/tests.TestLogger_Log_Level_Upper: This is an alert\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Upper: This is an error\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Upper: This is a warning\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Upper: This is a highlight\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Upper: This is an information\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Upper: This is a log\n" +
		"github.com/aliforever/go-log/tests.TestLogger_Log_Level_Upper: This is a trace\n"
	value := builder.String()

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}
}
