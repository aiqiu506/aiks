package aiks

import (
	"testing"
)

func TestDebugToFile(t *testing.T) {
	app := NewApp()
	app.SetConfigPath("config.yaml")
	app.InitComponent("log",)
	app.Start()
}
