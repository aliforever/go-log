package tests

import (
	"testing"

	"github.com/aliforever/go-log"
)

func TestLogger_LogPrefix(t *testing.T) {
	logger := log.NewLogger(nil)

	prefixedLogger := logger.Prefix("apple", "banana", "cherry")

	prefixedLogger.Log("Hello world!")

	doublePrefixedLogger := prefixedLogger.Prefix("date")

	doublePrefixedLogger.Log("I am here!")
}
