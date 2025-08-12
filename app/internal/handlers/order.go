package handlers

import (
	"demo/internal/models"
	"demo/internal/service"
	"demo/internal/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(or *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: or}
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
	err = oh.orderService.NewOrder(&order)
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

	start := time.Now()
	order, err := oh.orderService.GetOrderByUUID(uuid)
	elapsedTime := time.Now().Sub(start)
	fmt.Println("Elapsed time :%s", elapsedTime)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order with that id not found!"})
		return
	}
	orderDTO := models.OrderToDTO(order)
	c.JSON(http.StatusOK, orderDTO)

}
