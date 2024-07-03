package server

import (
	"context"
	"github.com/gofiber/fiber"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("server",
		fx.Provide(newConfig,
			NewServer,
		),
		fx.Invoke(func(lc fx.Lifecycle, cfg *Config, l *zap.Logger, app *fiber.App) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					l.Info("Start server")
					go app.Listen(cfg.Port)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return app.Shutdown()
				},
			})
		}),
	)
}
