package postgres

import (
	"WB/internal/postgres/models"
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

const (
	postgresMaxConn = 100
)

type DBRepository interface {
	InsertOrder(order models.Order) error
	GetAllOrders() ([]models.Order, error)
}

type PostgresDB struct {
	db *goqu.Database
}

type PostgresParams struct {
	fx.In

	Config *Config
}

func newPostgresDB(p PostgresParams) (*PostgresDB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.Config.Host,
		p.Config.Port,
		p.Config.User,
		p.Config.Password,
		p.Config.DBName,
		p.Config.SSLMode,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(postgresMaxConn)
	db.SetMaxIdleConns(postgresMaxConn)

	return &PostgresDB{
		db: goqu.New("postgres", db),
	}, nil
}

func (p *PostgresDB) InsertOrder(order models.Order) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	delivery := order.Delivery
	var deliveryId int
	_, err = tx.Insert("deliveries").
		Rows(
			delivery,
		).Returning("id").Executor().ScanVal(&deliveryId)

	if err != nil {
		return err
	}

	payment := order.Payment
	var paymentId int
	_, err = tx.Insert("payments").
		Rows(
			payment,
		).Returning("id").Executor().ScanVal(&paymentId)

	if err != nil {
		return err
	}

	items := order.Items
	//in my case len always equal 1
	var itemId int
	for _, item := range items {
		_, err = tx.Insert("items").
			Rows(
				item,
			).Returning("id").Executor().ScanVal(&itemId)

		if err != nil {
			return err
		}
	}

	_, err = tx.Insert("orders").
		Rows(goqu.Record{
			"order_uid":          order.OrderUid,
			"track_number":       order.TrackNumber,
			"entry":              order.Entry,
			"delivery_id":        deliveryId,
			"payment_id":         paymentId,
			"items_id":           itemId,
			"locale":             order.Locale,
			"internal_signature": order.InternalSignature,
			"customer_id":        order.CustomerId,
			"delivery_service":   order.DeliveryService,
			"shardkey":           order.Shardkey,
			"sm_id":              order.SmId,
			"date_created":       order.DateCreated,
			"oof_shard":          order.OofShard,
		},
		).Executor().Exec()

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order

	err := p.db.From(goqu.T("orders").As("o")).
		Join(
			goqu.T("deliveries").As("d"),
			goqu.On(goqu.I("d.id").Eq(
				goqu.I("o.delivery_id"),
			))).
		Join(
			goqu.T("payments").As("p"),
			goqu.On(goqu.I("p.id").Eq(
				goqu.I("o.payment_id"),
			))).Order(goqu.I("o.id").Asc()).ScanStructs(&orders)

	if err != nil {
		return nil, err
	}

	var items []models.Item

	err = p.db.From("items").Order(goqu.I("items.id").Asc()).ScanStructs(&items)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(orders); i++ {
		orders[i].Items = []models.Item{items[i]}
	}

	return orders, nil
}
