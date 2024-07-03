package order

import (
	in_memory "WB/internal/in-memory"
	"WB/internal/postgres"
	"WB/internal/postgres/models"
)

type UseCase interface {
	GetOrderByUID(orderUID string) (*models.Order, error)
	GetAllOrders() []*models.Order
}

type OrdersUC struct {
	cache in_memory.StorageRepository
	db    postgres.DBRepository
}

func newOrderUC(cache in_memory.StorageRepository, db postgres.DBRepository) *OrdersUC {
	return &OrdersUC{
		cache: cache,
		db:    db,
	}
}

func (uc *OrdersUC) GetOrderByUID(orderUID string) (*models.Order, error) {
	return uc.cache.GetByID(orderUID)
}

func (uc *OrdersUC) GetAllOrders() []*models.Order {
	return uc.cache.GetAll()
}
