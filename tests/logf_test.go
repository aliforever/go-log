package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aliforever/go-log"
)

func TestLogger_LogF(t *testing.T) {
	var builder strings.Builder
	logger := log.NewLogger(&builder).Level(6)

	var name string = "Ali"
	logger.LogF(fmt.Sprintf("My Name is %q", name))

	{
		var expected string = "github.com/aliforever/go-log/tests.TestLogger_LogF: My Name is \"Ali\"\n"
		var output string = builder.String()

		if expected != output {
			t.Errorf("The actual value is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", output)
		}
	}
}
