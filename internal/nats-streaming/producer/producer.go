package main

import (
	"WB/internal/postgres/models"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func main() {
	runProducer()
}

func runProducer() {
	sc, err := stan.Connect("test-cluster", "pub", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()
	defaultOrder := createDefaultOrder()

	data, err := json.Marshal(defaultOrder)
	if err != nil {
		log.Fatal(err)
	}
	sc.Publish("wbChan", data)
}

func createDefaultOrder() models.Order {
	delivery := models.Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	}

	payment := models.Payment{
		Transaction:  "b563feb7b2b84b6test",
		RequestId:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	}

	item := models.Item{
		ChrtId:      9934930,
		TrackNumber: "WBILMTESTTRACK",
		Price:       453,
		Rid:         "ab4219087a764ae0btest",
		Name:        "Mascaras",
		Sale:        30,
		Size:        "0",
		TotalPrice:  317,
		NmId:        2389212,
		Brand:       "Vivienne Sabo",
		Status:      202,
	}

	items := []models.Item{item}

	return models.Order{
		OrderUid:          uuid.New().String(),
		TrackNumber:       "WBILMTESTTRACK",
		Entry:             "WBIL",
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmId:              99,
		DateCreated:       time.Now(),
		OofShard:          "1",
	}
}
