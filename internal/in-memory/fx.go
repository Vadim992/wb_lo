package in_memory

import (
	"WB/internal/postgres"
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("cache",
		fx.Provide(
			fx.Annotate(
				newCache,
				fx.As(new(StorageRepository)),
			)),
		fx.Invoke(func(lc fx.Lifecycle, l *zap.Logger, repository StorageRepository, db postgres.DBRepository) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return recoverStorage(l, repository, db)
				},
			})
		}),
	)
}
