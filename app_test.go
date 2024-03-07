package clio


import (
	"testing"
)

func TestAppCreation(t *testing.T) {
	app := NewApp("test", "Application for test clio applications")
	if app.name != "test" {
		t.Error("Error on instating new App")
	}
	app.AddCmd("info", "Print information of application", infoHandlerTest)
	app.commands["info"].params["name"] = "name"
	if app.commands["info"].params["name"] != "name" {
		t.Error("Error to set a parameter in command.")
	}
	app.Run()
}

func infoHandlerTest(params Params, values Values) {
	print("testHandler")
}