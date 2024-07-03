package order

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module("OrderHandler",
		fx.Provide(
			fx.Annotate(
				newOrderUC,
				fx.As(new(UseCase)),
			),
			newOrderHandler,
		),
		fx.Invoke(MapRoutes),
	)
}
