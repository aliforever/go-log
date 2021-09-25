package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aliforever/go-log"
)

func TestLogger_LogPrefix(t *testing.T) {
	var builder = &strings.Builder{}

	logger := log.NewLogger(builder).Level(6)

	prefixedLogger := logger.Prefix("apple", "banana", "cherry")

	prefixedLogger.Log("Hello World!")

	expected := "github.com/aliforever/go-log/tests.TestLogger_LogPrefix: apple: banana: cherry: Hello World!\n"
	value := builder.String()

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}

	doublePrefixedLogger := prefixedLogger.Prefix("date")

	doublePrefixedLogger.Log("I am here!")

	expected = "github.com/aliforever/go-log/tests.TestLogger_LogPrefix: apple: banana: cherry: Hello World!\ngithub.com/aliforever/go-log/tests.TestLogger_LogPrefix: apple: banana: cherry: date: I am here!\n"
	value = builder.String()

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}
}
