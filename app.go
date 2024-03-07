package clio

import (
	"fmt"
	"os"
	"strings"
)


type App struct {
	name string
	description string
	commands map[string]*Command
}

func NewApp(name string, description string) App {
	app := App{
		name: name, 
		description: description,
		commands: map[string]*Command{},
	}
	createVersionCommand(&app)
	createHelpCommandApp(&app)
	for k, v := range app.commands {
		app.commands["help"].params[k] = v.description 
	}
	app.commands["help"].params["usage"] = app.name
	return app
}

func (app *App) AddCmd(
	name string,
	description string,
	handler handler,
) {
	cmd := NewCommand(name, description, handler)
	cmd.usage = fmt.Sprintf("%s %s", app.name, name)
	params := cmd.commands["help"].params
	params["usage"] = cmd.usage
	cmd.commands["help"].params = params
	app.commands[cmd.name] = cmd
	hlpCmd := app.commands["help"]
	hlpCmd.params[name] = description
}

func (app *App) Run() {
	args := os.Args
	for idx, arg := range args {
		if idx == 0 {
			continue
		}
		if strings.Index(arg, "--") == 0 {
			param := strings.Split(arg[2:], "=")
			key := param[0]
			value := param[1]
			print(key, value)
		} else {
			cmd := app.commands[arg]
			if cmd == nil {
				cmd = app.commands["help"]
				cmd.run([]string{})
			} else {
				cmd.run(args[idx + 1:])
			}
			break
		}
	}
}