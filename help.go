package clio

import "fmt"


func helpHandler(params Params, values Values) {
	name := params["name"]
	titleBlk := fmt.Sprintf("%s\n\n", name)
	print(titleBlk)
	usageBlk := fmt.Sprintf("Usage:\n\n\t\t%s <command> [arguments]\n\n", name)
	print(usageBlk)
	cmdBlk := "Commands:\n\n"
	for k, v := range params {
		if k == name {
			continue
		}
		prm := fmt.Sprintf("\t\t%s\t\t%s\n", k, v)
		cmdBlk = cmdBlk + prm
	}
	cmdBlk = cmdBlk + "\n\n"
	print(cmdBlk)
}


func createHelpCommand(app *App) {
	hlpCmd := NewCommand("help", helpHandler)
	hlpCmd.params["name"] = app.name
	for k, v := range app.commands {
		hlpCmd.params[k] = v.description
	}
}