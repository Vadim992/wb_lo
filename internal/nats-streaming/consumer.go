package nats_streaming

import (
	in_memory "WB/internal/in-memory"
	"WB/internal/postgres"
	"WB/internal/postgres/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"go.uber.org/zap"
	"log"
)

const url = "nats://%s:%s"

func ConnectToNatsStreaming(cfg *Config, ctx context.Context, l *zap.Logger, db postgres.DBRepository, cache in_memory.StorageRepository) {
	natsUrl := fmt.Sprintf(url, cfg.Host, cfg.Port)
	sc, err := stan.Connect(cfg.ClusterID, cfg.ClientID, stan.NatsURL(natsUrl))
	if err != nil {
		log.Fatal(err)
	}

	defer sc.Close()

	sc.Subscribe(cfg.WbChan, func(msg *stan.Msg) {
		l.Info(fmt.Sprint("get message from producer:", string(msg.Data)))

		var order models.Order

		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			l.Error(fmt.Sprint("err in subscribe decoder:", err))
		}

		err = db.InsertOrder(order)
		if err != nil {
			l.Error(fmt.Sprint("can't insert order in db:", err))
		}
		l.Info("Save data to db")
		cache.Save(order.OrderUid, &order)
		l.Info("Save data to cache")
	})

	select {
	case <-ctx.Done():
		l.Info("close connection with nats-streaming")
	}

}
