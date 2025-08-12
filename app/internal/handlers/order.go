package handlers

import (
	"demo/internal/models"
	"demo/internal/repository"
	"demo/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderRepository *repository.OrderRepository
}

func NewOrderHandler(or *repository.OrderRepository) *OrderHandler {
	return &OrderHandler{orderRepository: or}
}

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

func (oh *OrderHandler) GetOrderById(c *gin.Context) {
	uuid := c.Param("id")
	if len(uuid) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request, param id should be greater than 1"})
		return
	}

	order, err := oh.orderRepository.GetOrderByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order with that id not found!"})
		return
	}
	orderDTO := models.OrderToDTO(order)
	c.JSON(http.StatusOK, orderDTO)

}
