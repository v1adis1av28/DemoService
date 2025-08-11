package handlers

import (
	"demo/internal/models"
	"demo/internal/repository"
	"demo/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderRepository *repository.OrderRepository
}

func NewOrderHandler(or *repository.OrderRepository) *OrderHandler {
	return &OrderHandler{orderRepository: or}
}

// Получаем order из брокера, затем надо его проавалидировать
// После чего передать в слой репозитория и с помощью транзакции добавлять новую запись в таблици
func (oh *OrderHandler) HandleIncomingOrder(c *gin.Context) {
	var order models.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		log.Printf("Error while binding json %s", err.Error())
		return
	}
	if !utils.ValidateOrder(&order) {
		log.Printf("Error on validating order with id %s", order.OrderUID)
		return
	}
	err = oh.orderRepository.NewOrder(&order)
	if err != nil {
		log.Printf("Error on creating new order %s", err.Error())
		return
	}

	log.Printf("Order was succesfully created, order_id: %s", order.OrderUID)

}
