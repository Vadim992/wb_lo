package nats_streaming

import (
	in_memory "WB/internal/in-memory"
	"WB/internal/postgres"
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"nats-streaming",
		fx.Provide(
			newConfig,
		),
		fx.Invoke(func(lc fx.Lifecycle, cfg *Config, l *zap.Logger, db postgres.DBRepository, cache in_memory.StorageRepository) {
			ctx, cancell := context.WithCancel(context.Background())

			lc.Append(fx.Hook{
				OnStart: func(c context.Context) error {
					go ConnectToNatsStreaming(cfg, ctx, l, db, cache)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					cancell()
					return nil
				},
			})

		}),
	)
}
