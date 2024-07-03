package main

import (
	"WB/config"
	in_memory "WB/internal/in-memory"
	nats_streaming "WB/internal/nats-streaming"
	"WB/internal/order"
	"WB/internal/postgres"
	"WB/internal/server"
	"WB/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.NewModule(),
		logger.NewModule(),
		in_memory.NewModule(),
		postgres.NewModule(),
		nats_streaming.NewModule(),
		server.NewModule(),
		order.NewModule(),
	).Run()
}
