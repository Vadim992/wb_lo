package order

import "github.com/gofiber/fiber"

const (
	staticPath = "./static"
)

func MapRoutes(r *fiber.App, h *OrderHandler) {
	r.Static("/orders", staticPath)
	r.Post("/orders", h.PostOrders())

	r.Get("/orders/all", h.AllOrdersHandler())

	r.Get("/orders/:order_uid", h.OrderIDHandler())
}
