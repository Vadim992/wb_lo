package postgres

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module("postgres",
		fx.Provide(
			newConfig,
			fx.Annotate(
				newPostgresDB,
				fx.As(new(DBRepository))),
		),
	)
}
