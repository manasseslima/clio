package clio


import (
	"strings"
)

type Param struct {
	Name string 
	Description string
	Type string
	Required bool
	Value any
}

func NewParam(name string, description string, required bool) Param {
	param := Param{
		Name: name,
		Description: description,
		Type: "str",
		Required: required,
		Value: nil,
	}
	return param
}

type Params map[string]Param
type Values []string
type handler func(Params, Values)


type Command struct {
	Name string
	Description string
	Usage string
	handler handler
	Commands map[string]*Command
	Params Params
	Values Values
}

func NewCommand(name string, description string, handler handler) *Command {
	cmd := newCommandWithoutHelp(name, description, handler)
	cmd.Usage = name
	createHelpCommandCmd(cmd)
	for k, v := range cmd.Commands {
		param := NewParam(k, v.Description, false)
		param.Value = v.Description
		cmd.Commands["help"].Params[k] = param
	}
	return cmd
}

func newCommandWithoutHelp(name string, description string, handler handler) *Command {
	cmd := &Command{
		Name: name, 
		Description: description,
		handler: handler,
		Commands: map[string]*Command{},
		Params: Params{},
		Values: []string{},
	}
	return cmd
}

func (cmd *Command) NewParam(name string, description string, required bool, typee string) {
	param := Param{
		Name: name,
		Description: description,
		Type: typee,
		Required: required,
		Value: nil,
	}
	cmd.Params[param.Name] = param
}

func (cmd *Command) AddParam(param Param) {
	cmd.Params[param.Name] = param
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
			splt := strings.Split(arg[2:], "=")
			key := splt[0]
			value := splt[1]
			param := cmd.Params[key]
			param.Value = value
			cmd.Params[key] = param
		} else {
			cmd.Values = append(cmd.Values, arg)
		}
	}
	if runme {
		cmd.handler(cmd.Params, cmd.Values)
	}
}