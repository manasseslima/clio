package clio

import (
	"testing"
)

func infoHandlerTest(params Params, values Values) {
	print("testHandler")
}

func createNewApp() *App {
	app := NewApp("test", "Application for test clio applications")
	app.NewCmd("info", "Print information of application", infoHandlerTest)
	param := NewParam("name", "Name of param", false)
	param.Value = "name"
	app.Commands["info"].AddParam(param)
	return &app
}

func TestAppCreation(t *testing.T) {
	app := createNewApp()
	if app.Name != "test" {
		t.Error("Error on instating new App")
	}
	if app.Commands["info"].Params["name"].Value != "name" {
		t.Error("Error to set a parameter in command.")
	}
	app.Run()
}

func TestAppGetCmd(t *testing.T) {
	app := createNewApp()
	cmd := app.GetCmd("info help")
	if cmd == nil {
		t.Error("Not find command by method GetCmd().")
	}
}
