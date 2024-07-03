package in_memory

import (
	"WB/internal/postgres"
	"WB/internal/postgres/models"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

type StorageRepository interface {
	GetByID(id string) (*models.Order, error)
	GetAll() []*models.Order
	Save(id string, order *models.Order)
}

type Cache struct {
	mu    sync.Mutex
	cache map[string]*models.Order
}

func newCache() *Cache {
	return &Cache{
		cache: make(map[string]*models.Order, 0),
	}
}

func (c *Cache) GetByID(id string) (*models.Order, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	order, ok := c.cache[id]
	if !ok {
		return nil, noExistIdErr
	}

	return order, nil
}

func (c *Cache) GetAll() []*models.Order {
	c.mu.Lock()
	defer c.mu.Unlock()

	orders := make([]*models.Order, 0, len(c.cache))
	for _, order := range c.cache {
		orders = append(orders, order)
	}

	return orders
}

func (c *Cache) Save(id string, order *models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[id] = order
}

func recoverStorage(l *zap.Logger, storage StorageRepository, db postgres.DBRepository) error {
	orders, err := db.GetAllOrders()

	if err != nil {
		l.Error(fmt.Sprint("recover storage err: ", err))
		return err
	}

	for _, order := range orders {
		copyOrder := order
		storage.Save(order.OrderUid, &copyOrder)
	}

	return nil
}
