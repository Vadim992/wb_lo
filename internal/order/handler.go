package order

import (
	"github.com/gofiber/fiber"
	"go.uber.org/zap"
)

const (
	formOrderUID  = "order_uid"
	paramOrderUID = "order_uid"
)

type ErrorResult struct {
	ErrMsg string `json:"err_msg"`
}

type OrderHandler struct {
	uc UseCase
	l  *zap.Logger
}

func newOrderHandler(uc UseCase, l *zap.Logger) *OrderHandler {
	return &OrderHandler{
		uc: uc,
		l:  l,
	}
}

func (h *OrderHandler) OrderIDHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) {
		orderUID := ctx.Params(paramOrderUID)
		if orderUID == "" {
			ctx.Status(fiber.StatusBadRequest)
			return
		}

		order, err := h.uc.GetOrderByUID(orderUID)
		if err != nil {
			ctx.Status(fiber.StatusNotFound).JSON(
				ErrorResult{ErrMsg: err.Error()},
			)
			return
		}

		ctx.JSON(order)
	}
}

func (h *OrderHandler) AllOrdersHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) {
		orders := h.uc.GetAllOrders()

		ctx.JSON(orders)
	}
}

func (h *OrderHandler) PostOrders() fiber.Handler {
	return func(ctx *fiber.Ctx) {
		formValue := ctx.FormValue(formOrderUID)
		if formValue != "" {
			ctx.Redirect("./" + formValue)
		} else {
			ctx.Redirect("./")
		}
	}
}
