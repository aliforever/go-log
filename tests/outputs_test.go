package tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/aliforever/go-log"
)

func TestLogger_Output_String_Builder(t *testing.T) {
	var builder strings.Builder

	logger := log.NewLogger(&builder).Level(6)

	logger.Log("Testing String Builder", "Done")

	expected := "github.com/aliforever/go-log/tests.TestLogger_Output_String_Builder: Testing String Builder Done\n"
	value := builder.String()

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}
}

func TestLogger_Output_File(t *testing.T) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 644)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	logger := log.NewLogger(file).Level(6)

	logger.Log("Testing File Log", "Done")

	data, err := ioutil.ReadFile("log.txt")
	if err != nil {
		t.Fatal(err)
	}

	expected := "github.com/aliforever/go-log/tests.TestLogger_Output_File: Testing File Log Done\n"
	value := string(data)

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nValue: %s", expected, value))
	}
}
