package delivery

import (
	"net/http"
	"order-service/internal/domain"
	"order-service/internal/dto"
	"order-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService *usecase.OrderService
}

func NewOrderController(orderService *usecase.OrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	userID, _ := c.Get("userID")

	var request dto.CreateOrderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := &domain.Order{
		UserID:   userID.(uint),
		Products: request.Products,
	}

	response, err := oc.OrderService.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (oc *OrderController) GetOrders(c *gin.Context) {
	userID, _ := c.Get("userID")

	orders, err := oc.OrderService.GetOrders(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
