package clio

import "fmt"


func versionHandler(params Params, values Values) {
	name := params["name"]
	version := params["version"]
	ver := fmt.Sprintf("%s %s\n", name, version)
	print(ver)
}

func createVersionCommand(app *App) {
	verCmd := NewCommand("version", versionHandler)
	verCmd.params["name"] = app.name
	verCmd.params["version"] = "v0.0.0"
	app.commands[verCmd.name] = verCmd
}