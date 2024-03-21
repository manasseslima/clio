package clio


import (
	"strings"
)


type Params map[string]string
type Values []string
type handler func(Params, Values)


type Command struct {
	Name string
	Description string
	Usage string
	handler handler
	Commands map[string]*Command
	Params map[string]string
	Values []string
}

func NewCommand(name string, description string, handler handler) *Command {
	cmd := newCommandWithoutHelp(name, description, handler)
	cmd.Usage = name
	createHelpCommandCmd(cmd)
	for k, v := range cmd.Commands {
		cmd.Commands["help"].Params[k] = v.Description 
	}
	return cmd
}

func newCommandWithoutHelp(name string, description string, handler handler) *Command {
	cmd := &Command{
		Name: name, 
		Description: description,
		handler: handler,
		Commands: map[string]*Command{},
		Params: map[string]string{},
		Values: []string{},
	}
	return cmd
}

func (cmd *Command) run(args []string) {
	runme := true
	for idx, arg := range args {
		if idx == 0 && !strings.Contains(arg, "--") {
			cd := cmd.Commands[arg]
			if cd == nil {
				cd = cmd.Commands["help"]
				cd.run([]string{})
			} else {
				cd.run(args[idx + 1:])
			}
			runme = false
			break
		}
		if strings.Index(arg, "--") == 0 {
			param := strings.Split(arg[2:], "=")
			key := param[0]
			value := param[1]
			cmd.Params[key] = value
		} else {
			cmd.Values = append(cmd.Values, arg)
		}
	}
	if runme {
		cmd.handler(cmd.Params, cmd.Values)
	}
}