package clio

import (
	"fmt"
	"os"
	"strings"
)


type App struct {
	Name string
	Description string
	Commands map[string]*Command
}

func NewApp(name string, description string) App {
	app := App{
		Name: name, 
		Description: description,
		Commands: map[string]*Command{},
	}
	createVersionCommand(&app)
	createHelpCommandApp(&app)
	for k, v := range app.Commands {
		app.Commands["help"].Params[k] = v.Description 
	}
	app.Commands["help"].Params["usage"] = app.Name
	return app
}

func (app *App) NewCmd(
	name string,
	description string,
	handler handler,
) {
	cmd := NewCommand(name, description, handler)
	cmd.Usage = fmt.Sprintf("%s %s", app.Name, name)
	params := cmd.Commands["help"].Params
	params["usage"] = cmd.Usage
	cmd.Commands["help"].Params = params
	app.Commands[cmd.Name] = cmd
	hlpCmd := app.Commands["help"]
	hlpCmd.Params[name] = description
}

func (app *App) AddCmd(cmd *Command) {
	app.Commands[cmd.Name] = cmd
)

func (app *App) GetCmd(path string) *Command {
	spt := strings.Split(path, " ")
	cmds :=  app.Commands
	var cur *Command
	for idx, scmd := range spt {
		cur = cmds[scmd]
		if cur == nil {
			return nil
		}
		if idx != len(spt) - 1 {
			cmds = cur.Commands
		}
	}
	return cur
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
			cmd := app.Commands[arg]
			if cmd == nil {
				cmd = app.Commands["help"]
				cmd.run([]string{})
			} else {
				cmd.run(args[idx + 1:])
			}
			break
		}
	}
}