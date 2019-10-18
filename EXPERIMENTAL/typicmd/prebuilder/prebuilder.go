package prebuilder

import (
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typicmd/prebuilder/walker"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typienv"
)

type prebuilder struct {
	Annotated     *AnnotatedGenerator
	Configuration *ConfigurationGenerator
	TestTarget    *TestTargetGenerator
}

func (p *prebuilder) Initiate(ctx *typictx.Context) (err error) {
	log.Debug("Scan project to get package and filenames")
	root := typienv.AppName
	packages, filenames, err := scanProject(root)
	if err != nil {
		return
	}
	log.Debug("Walk the project to get annotated or metadata")
	projectFiles, err := walker.WalkProject(filenames)
	if err != nil {
		return
	}
	log.Debug("Walk the context file")
	contextFile, err := walker.WalkContext(ctxPath)
	if err != nil {
		return
	}
	p.Annotated = &AnnotatedGenerator{
		Root:         ctx.Root,
		ProjectFiles: projectFiles,
		Packages:     packages,
	}
	p.Configuration = &ConfigurationGenerator{
		Configs:     createConfigs(ctx),
		ContextFile: contextFile,
	}
	p.TestTarget = &TestTargetGenerator{
		Root:     ctx.Root,
		Packages: packages,
	}
	return
}

func (p *prebuilder) Prebuild() (r report, err error) {
	if r.TestTargetUpdated, err = p.TestTarget.Generate(); err != nil {
		return
	}
	if r.AnnotatedUpdated, err = p.Annotated.Generate(); err != nil {
		return
	}
	if r.ConfigurationUpdated, err = p.Configuration.Generate(); err != nil {
		return
	}
	return
}

func createConfigs(ctx *typictx.Context) (configs []Config) {
	configs = append(configs, Config{
		Key: strcase.ToCamel(strings.ToLower(ctx.Application.Prefix)),
		Typ: reflect.TypeOf(ctx.Application.Spec).String(),
	})
	for _, module := range ctx.Modules {
		configs = append(configs, Config{
			Key: strcase.ToCamel(strings.ToLower(module.Prefix)),
			Typ: reflect.TypeOf(module.Spec).String(),
		})
	}
	return
}
