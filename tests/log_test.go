package tests

import (
	"strings"
	"testing"

	"github.com/aliforever/go-log"
)

func TestLogger_Log(t *testing.T) {
	var output strings.Builder

	logger := log.NewLogger(&output).Level(6)

	var name string = "Cherry"
	logger.Log("Fruit name is", name)

	{
		var expected string = "github.com/aliforever/go-log/tests.TestLogger_Log: Fruit name is Cherry\n"
		var actual string = output.String()

		if expected != actual {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
		}
	}
}
