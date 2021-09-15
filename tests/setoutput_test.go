package tests

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/aliforever/go-log"
)

func TestLogger_SetOutput_String_Builder(t *testing.T) {
	logger := log.NewLogger()

	var builder strings.Builder

	logger.SetOutput(&builder)

	logger.Log("Testing String Builder", "Done")

	logger.SetOutput(os.Stdout)
	logger.LogF("%s", builder.String())
}

func TestLogger_SetOutput_File(t *testing.T) {
	logger := log.NewLogger()

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 644)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	logger.SetOutput(file)
	logger.Log("Testing File Log", "Done")

	data, err := ioutil.ReadFile("log.txt")
	if err != nil {
		t.Fatal(err)
	}

	logger.SetOutput(os.Stdout)
	logger.LogF("%s", data)
}

func TestLogger_SetOutput_Stdout(t *testing.T) {
	logger := log.NewLogger()
	logger.SetOutput(os.Stdout)

	var name string = "Ali"
	logger.LogF("My name is %q", name)
}
