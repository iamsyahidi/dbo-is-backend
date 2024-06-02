package controllers

import (
	"dbo/dbo-is-backend/middleware"
	"dbo/dbo-is-backend/models"
	"dbo/dbo-is-backend/pkg/logger"
	"dbo/dbo-is-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderService services.OrderServiceInterface
}

type OrderControllerInterface interface {
	CreateOrder(c *gin.Context)
}

func NewOrderController(orderService services.OrderServiceInterface) OrderControllerInterface {
	return &orderController{
		orderService: orderService,
	}
}

// CreateOrder godoc
// @Summary Create an order
// @Description Create an order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body models.OrderRegister true "Order"
// @Response 201 {object} models.Response
// @Response 500 {object} models.Response
// @Response 400 {object} models.Response
// @Response 404 {object} models.Response
// @Router /orders [post]
func (oc *orderController) CreateOrder(c *gin.Context) {
	v, ok := c.Get("customer")
	if !ok {
		c.JSON(401, models.Response{
			Code:    http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	var orderRegister models.OrderRegister
	if err := c.ShouldBindJSON(&orderRegister); err != nil {
		middleware.Response(c, orderRegister, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	order := models.Order{
		Payment:    orderRegister.Payment,
		CustomerID: v.(*models.CustomerClaims).ID,
		CreatedBy:  v.(*models.CustomerClaims).Name,
	}

	var orderDetail []models.OrderDetail
	for _, e := range orderRegister.Products {
		detail := models.OrderDetail{
			ProductID: e.ProductID,
			Qty:       e.Qty,
		}
		orderDetail = append(orderDetail, detail)
	}

	response, err := oc.orderService.CreateOrder(&order, &orderDetail)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, orderRegister, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, orderRegister, *response)
}
