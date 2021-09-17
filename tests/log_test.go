package tests

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/aliforever/go-log"
)

func TestLogger_Log_Begin_End(t *testing.T) {
	logger := log.NewLogger(nil)

	localLogger := logger.Begin()

	var name string = "Apple"
	localLogger.LogF("The name is %q", name)
	fmt.Printf("Hello %s!\n", name)

	time.Sleep(time.Microsecond * 2)
	localLogger.End()
}

func TestLogger_Log_Apple(t *testing.T) {
	logger := log.NewLogger(nil)

	var name string = "Apple"
	logger.Log("Fruit name is", name)
}

func TestLogger_Log_Banana(t *testing.T) {
	logger := log.NewLogger(nil)

	var name string = "Banana"
	logger.Log("Fruit name is", name)
}

func TestLogger_Log_Cherry(t *testing.T) {
	var output strings.Builder

	logger := log.NewLogger(&output)

	var name string = "Cherry"
	logger.Log("Fruit name is", name)

	{
		var expected string = "Fruit name is Cherry\n"
		var actual string = output.String()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
		}
	}
}
