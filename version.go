package clio

import "fmt"


func versionHandler(params Params, values Values) {
	name := params["name"].Value
	version := params["version"].Value
	ver := fmt.Sprintf("%s %s\n", name, version)
	print(ver)
}

func createVersionCommand(app *App) {
	verCmd := NewCommand("version", "Print application version", versionHandler)
	var param Param
	param = NewParam("name", "name", false)
	param.Value = app.Name
	verCmd.Params["name"] = param
	param = NewParam("version", "version", false)
	param.Value = "v0.0.0"
	verCmd.Params["version"] = param
	app.Commands[verCmd.Name] = verCmd
}