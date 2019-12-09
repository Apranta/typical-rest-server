package typserver

import (
	"context"
	"fmt"
	"time"

	logrusmiddleware "github.com/bakatz/echo-logrusmiddleware"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

// Config is server configuration
type Config struct {
	Debug bool `default:"false"`
}

// Module of Server
type Module struct{}

// Configure server
func (s *Module) Configure() (prefix string, spec, loadFn interface{}) {
	prefix = "SERVER"
	spec = &Config{}
	loadFn = func(loader typcfg.Loader) (cfg Config, err error) {
		err = loader.Load(prefix, &cfg)
		return
	}
	return
}

// Provide dependencies
func (s *Module) Provide() []interface{} {
	return []interface{}{
		s.Create,
	}
}

// Destroy dependencies
func (s *Module) Destroy() []interface{} {
	return []interface{}{
		s.Shutdown,
	}
}

// Create new server
func (s *Module) Create(cfg Config) *echo.Echo {
	server := echo.New()
	server.HideBanner = true
	server.Debug = cfg.Debug
	logrusMwConfig := logrusmiddleware.Config{}
	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
		logrusMwConfig.IncludeRequestBodies = true
		logrusMwConfig.IncludeResponseBodies = true
	} else {
		log.SetLevel(log.WarnLevel)
		log.SetFormatter(&log.JSONFormatter{})
	}
	server.Use(logrusmiddleware.HookWithConfig(logrusMwConfig))
	server.Logger = logrusmiddleware.Logger{Logger: log.StandardLogger()}
	return server
}

// Shutdown the server
func (s *Module) Shutdown(server *echo.Echo) error {
	fmt.Println("Server is shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}
