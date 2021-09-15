package tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/aliforever/go-log"
)

func TestLogger_SetOutput_String_Builder(t *testing.T) {
	logger := log.NewLogger(nil)

	var builder strings.Builder

	logger.Log("Testing String Builder", "Done")

	logger.LogF("%s", builder.String())
}

func TestLogger_SetOutput_File(t *testing.T) {
	logger := log.NewLogger(nil)

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 644)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	logger.Log("Testing File Log", "Done")

	data, err := ioutil.ReadFile("log.txt")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(data))
}

func TestLogger_SetOutput_Stdout(t *testing.T) {
	logger := log.NewLogger(nil)

	var name string = "Ali"
	logger.LogF("My name is %q", name)
}
