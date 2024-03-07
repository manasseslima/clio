package clio

import "fmt"


func helpHandler(params Params, values Values) {
	usage := params["usage"]
	titleBlk := fmt.Sprintf("%s\n\n", usage)
	print(titleBlk)
	usageBlk := fmt.Sprintf("Usage:\n\n\t%s <command> [arguments]\n\n", usage)
	print(usageBlk)
	cmdBlk := "Commands:\n\n"
	for k, v := range params {
		if k == "name" || k == "usage" {
			continue
		}
		prm := fmt.Sprintf("\t%s\t\t%s\n", k, v)
		cmdBlk = cmdBlk + prm
	}
	cmdBlk = cmdBlk + "\n\n"
	print(cmdBlk)
}


func createHelpCommandApp(app *App) {
	hlpCmd := newCommandWithoutHelp("help", "Print this help text", helpHandler)
	hlpCmd.params["name"] = app.name
	app.commands["help"] = hlpCmd
}

func createHelpCommandCmd(cmd *Command) {
	hlpCmd := newCommandWithoutHelp("help", "Print this help text", helpHandler)
	hlpCmd.params["name"] = cmd.name
	hlpCmd.params["usage"] = cmd.usage
	cmd.commands["help"] = hlpCmd
}