package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-log"
)

func TestLogger_LogF(t *testing.T) {
	logger := log.NewLogger(nil)

	var name string = "Ali"
	logger.LogF(fmt.Sprintf("My Name is %q", name))
}
