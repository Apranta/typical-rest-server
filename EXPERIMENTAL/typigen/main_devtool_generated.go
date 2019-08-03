package typigen

import (
	"os"

	"github.com/typical-go/runn"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/internal/bash"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typienv"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typigen/generated"

	log "github.com/sirupsen/logrus"
)

// MainDevToolGenerated to generate code in typical package
func MainDevToolGenerated(t typictx.Context) (err error) {
	filename := typienv.TypicalDevToolMainPackage() + "/generated.go"

	recipe := generated.SourceRecipe{
		PackageName: "main",
	}

	for _, lib := range devtoolSideEffects(t) {
		recipe.AddImportPogo(generated.ImportPogo{Alias: "_", PackageName: lib})
	}

	if recipe.Blank() {
		os.Remove(filename)
		return
	}

	log.Infof("Generate recipe for Typical-Dev-Tool: %s", filename)
	return runn.Execute(
		recipe.Cook(filename),
		bash.GoFmt(filename),
	)
}

func devtoolSideEffects(t typictx.Context) (sideEffects []string) {
	for _, module := range t.Modules {
		for _, sideEffect := range module.SideEffects {
			if sideEffect.TypicalDevToolFlag {
				sideEffects = append(sideEffects, sideEffect.Library)
			}
		}
	}
	return
}