package clio

import "fmt"


func versionHandler(params Params, values Values) {
	name := params["name"]
	version := params["version"]
	ver := fmt.Sprintf("%s %s\n", name, version)
	print(ver)
}

func createVersionCommand(app *App) {
	verCmd := NewCommand("version", "Print application version", versionHandler)
	verCmd.Params["name"] = app.Name
	verCmd.Params["version"] = "v0.0.0"
	app.Commands[verCmd.Name] = verCmd
}