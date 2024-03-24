package clio

import "fmt"


func helpHandler(params Params, values Values) {
	usage := params["usage"].Value
	titleBlk := fmt.Sprintf("%s\n\n", usage)
	print(titleBlk)
	usageBlk := fmt.Sprintf("Usage:\n\n\t%s <command> [arguments]\n\n", usage)
	print(usageBlk)
	cmdBlk := "Commands:\n\n"
	for k, v := range params {
		if k == "name" || k == "usage" {
			continue
		}
		prm := fmt.Sprintf("\t%s\t\t%s\n", k, v.Value)
		cmdBlk = cmdBlk + prm
	}
	cmdBlk = cmdBlk + "\n\n"
	print(cmdBlk)
}


func createHelpCommandApp(app *App) {
	hlpCmd := newCommandWithoutHelp("help", "Print this help text", helpHandler)
	param := NewParam("name", "name", false)
	param.Value =  app.Name
	hlpCmd.Params["name"] = param
	app.Commands["help"] = hlpCmd
}

func createHelpCommandCmd(cmd *Command) {
	hlpCmd := newCommandWithoutHelp("help", "Print this help text", helpHandler)
	var param Param
	param = hlpCmd.Params["name"]
	param.Value = cmd.Name
	hlpCmd.Params["name"] = param
	param = hlpCmd.Params["usage"]
	param.Value = cmd.Usage
	hlpCmd.Params["usage"] = param
	cmd.Commands["help"] = hlpCmd
}