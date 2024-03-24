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
		param := NewParam(k, v.Description, false)
		param.Value = v.Description
		app.Commands["help"].Params[k] = param
	}
	param := NewParam("usage", "Usage", false)
	param.Value = app.Name
	app.Commands["help"].AddParam(param)
	return app
}

func (app *App) NewCmd(
	name string,
	description string,
	handler handler,
) {
	cmd := NewCommand(name, description, handler)
	app.addHelpCommand(cmd)
}

func (app *App) addHelpCommand(cmd *Command) {
	cmd.Usage = fmt.Sprintf("%s %s", app.Name, cmd.Name)
	var param Param
	param = NewParam("usage", "usage", false)
	param.Value = cmd.Usage
	cmd.Commands["help"].AddParam(param)
	app.Commands[cmd.Name] = cmd
	param = NewParam(cmd.Name, cmd.Name, false)
	param.Value = cmd.Description
	app.Commands["help"].AddParam(param)
}

func (app *App) AddCmd(cmd *Command) {
	app.addHelpCommand(cmd)
}

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