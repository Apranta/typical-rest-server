package buildtool

import (
	"fmt"
	"io"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typicmd/buildtool/markdown"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typiobj"
)

func readme(w io.Writer, ctx *typictx.Context) (err error) {
	md := &markdown.Markdown{Writer: w}
	md.Comment("Autogenerated by Typical-Go. DO NOT EDIT.")
	if ctx.Name != "" {
		md.H1(ctx.Name)
	} else {
		md.H1("Typical-Go Project")
	}
	if ctx.Description != "" {
		md.Writeln(ctx.Description)
	}
	md.H2("Prerequisite")
	prerequisite(md)
	md.H2("Infrastructure")
	infrastructure(md)
	md.H2("Run & Test")
	runInstruction(md)
	md.H2("Application")
	application(md, ctx.Application)
	md.H2("Modules")
	for _, m := range ctx.Modules {
		module(md, m)
	}
	md.H2("Release Distribution")
	releaseDistribution(md)
	return
}

func prerequisite(md *markdown.Markdown) {
	md.OrderedList(
		"[Go](https://golang.org/doc/install) (It is recommend to install via [Homebrew](https://brew.sh/) `brew install go`)",
		"[Docker Compose](https://docs.docker.com/compose/install/)",
	)
}

func infrastructure(md *markdown.Markdown) {
	md.Writeln("Use `./typicalw docker up` to spin up infrastructure docker.")
	md.Writeln("Use `./typicalw docker compose` to generate \"docker-compose.yml\" based on Typical Context.")
}

func runInstruction(md *markdown.Markdown) {
	md.Writeln("Use `./typicalw run` to compile and run local development.")
	md.Writeln("Use `./typicalw test` to execute the unit testing.")
	md.Writeln("[Learn more](https://typical-go.github.io/learn-more/wrapper.html)")
}

func releaseDistribution(md *markdown.Markdown) {
	md.Writeln("Use `./typicalw release` to make the release. You can find the binary at `release` folder.")
	md.Writeln("[Learn more](https://typical-go.github.io/learn-more/release.html)")
}

func application(md *markdown.Markdown, app interface{}) {
	if configurer, ok := app.(typiobj.Configurer); ok {
		configTable(md, configurer.Configure().ConfigFields())
	}
}

func module(md *markdown.Markdown, module interface{}) {
	if name := typiobj.Name(module); name != "" {
		md.H3(strcase.ToCamel(name))
	}
	if description := typiobj.Description(module); description != "" {
		md.Writeln(description)
	}
	if configurer, ok := module.(typiobj.Configurer); ok {
		configTable(md, configurer.Configure().ConfigFields())
	}
	if cli, ok := module.(typiobj.BuildCLI); ok {
		md.WriteString("Commands:\n")
		cmd := cli.Command()
		var cmdHelps []string
		for _, subcmd := range cmd.Subcommands {
			cmdHelps = append(cmdHelps, fmt.Sprintf("`./typicalw %s %s`: %s", cmd.Name, subcmd.Name, subcmd.Usage))
		}
		md.UnorderedList(cmdHelps...)
	}
}

func configTable(md *markdown.Markdown, fields []typiobj.ConfigField) {
	md.WriteString("| Name | Type | Default | Required |\n")
	md.WriteString("|---|---|---|:---:|\n")
	for _, field := range fields {
		var required string
		if field.Required {
			required = "Yes"
		}
		md.WriteString(fmt.Sprintf("|%s|%s|%s|%s|\n",
			field.Name, field.Type, field.Default, required))
	}
	md.WriteString("\n")
}
