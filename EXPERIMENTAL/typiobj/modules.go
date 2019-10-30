package typiobj

import (
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/slice"
	"github.com/urfave/cli"
	"go.uber.org/dig"
)

// Modules is list of module
type Modules slice.Interfaces

// Configurations return list of config
func (m Modules) Configurations() (cfgs []Configuration) {
	for _, module := range m {
		if configurer, ok := module.(Configurer); ok {
			cfgs = append(cfgs, configurer.Configure())
		}
	}
	return
}

// Commands return list of command
func (m Modules) Commands() (cmds []cli.Command) {
	for _, module := range m {
		if cmdline, ok := module.(CommandLiner); ok {
			cmds = append(cmds, cmdline.CommandLine())
		}
	}
	return
}

// Provide dependency
func (m Modules) Provide() (constructors []interface{}) {
	for _, module := range m {
		if provider, ok := module.(Provider); ok {
			constructors = append(constructors, provider.Provide()...)
		}
	}
	return
}

// Destruct dependency
func (m Modules) Destruct(c *dig.Container) (err error) {
	for _, module := range m {
		if destructor, ok := module.(Destructor); ok {
			if err = destructor.Destruct(c); err != nil {
				return
			}
		}
	}
	return
}

// Helps information
func (m Modules) Helps() (helps []Help) {
	for _, module := range m {
		var help Help
		help.Name = Name(module)
		help.Description = Description(module)
		// TODO: configuration
		helps = append(helps, help)
	}
	return
}