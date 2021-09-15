package tests

import (
	"testing"

	"github.com/aliforever/go-log"
)

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
	logger := log.NewLogger(nil)

	var name string = "Cherry"
	logger.Log("Fruit name is", name)
}
