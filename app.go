package clio

import (
	"os"
	"strings"
)


type App struct {
	name string
	commands map[string]*Command
}

func NewApp(name string) App {
	app := App{name: name, commands: map[string]*Command{} }
	createVersionCommand(&app)
	createHelpCommand(&app)
	return app
}

func (app *App) AddCmd(
	name string,
	handler handler,
) {
	cmd := NewCommand(name, handler)
	app.commands[cmd.name] = cmd
}

func (app *App) Run() {
	for idx, arg := range os.Args {
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
			cmd.run(os.Args[idx + 1:])
			break
		}
	}
}