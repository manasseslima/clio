package clio


import (
	"strings"
)


type Params map[string]string
type Values []string
type handler func(Params, Values)


type Command struct {
	name string
	description string
	usage string
	handler handler
	commands map[string]*Command
	params map[string]string
	values []string
}

func NewCommand(name string, description string, handler handler) *Command {
	cmd := newCommandWithoutHelp(name, description, handler)
	cmd.usage = name
	createHelpCommandCmd(cmd)
	for k, v := range cmd.commands {
		cmd.commands["help"].params[k] = v.description 
	}
	return cmd
}

func newCommandWithoutHelp(name string, description string, handler handler) *Command {
	cmd := &Command{
		name: name, 
		description: description,
		handler: handler,
		commands: map[string]*Command{},
		params: map[string]string{},
		values: []string{},
	}
	return cmd
}

func (cmd *Command) run(args []string) {
	runme := true
	for idx, arg := range args {
		if idx == 0 && !strings.Contains(arg, "--") {
			cmd := cmd.commands[arg]
			if cmd == nil {
				cmd = cmd.commands["help"]
				cmd.run([]string{})
			} else {
				cmd.run(args[idx + 1:])
			}
			
			runme = false
			break
		}
		if strings.Index(arg, "--") == 0 {
			param := strings.Split(arg[2:], "=")
			key := param[0]
			value := param[1]
			cmd.params[key] = value
		} else {
			cmd.values = append(cmd.values, arg)
		}
	}
	if runme {
		cmd.handler(cmd.params, cmd.values)
	}
}