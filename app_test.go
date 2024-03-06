package clio


import (
	"testing"
)

func TestAppCreation(t *testing.T) {
	app := NewApp("test")
	if app.name != "test" {
		t.Error("Error on instating App")
	}
	app.AddCmd("info", infoHandlerTest)
	app.commands["info"].params["name"] = "name"
}

func infoHandlerTest(params Params, values Values) {
	print("testHandler")
}