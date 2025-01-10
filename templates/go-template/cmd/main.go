package main

import (
	"github.com/litmus-zhang/go-template/internal/api"
	"github.com/litmus-zhang/go-template/internal/config"
	"github.com/litmus-zhang/go-template/internal/db"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		config.Module,
		db.Module,
		api.Module,
		fx.Provide(zap.NewDevelopment),
		fx.Invoke(func(cfg *config.Config, server *api.Server) error {
			return server.Start()
		}),
	).Run()
}
