package controllers

import (
	"encoding/json"
	"go-restapi/params"
	"go-restapi/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *services.OrderServices
}

func NewOrderController(orderService *services.OrderServices) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

// func (o *OrderController) CreateOrder(rw http.ResponseWriter, r *http.Request) {
func (o *OrderController) CreateOrder(ctx *gin.Context) {
	var req params.OrderCreate

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, &response)
		return
	}

	response := o.orderService.CreateOrder(&req)
	params.WriteJsonResponse(ctx.Writer, response)
}

// func (o *OrderController) GetAllOrders(rw http.ResponseWriter, r *http.Request) {
func (o *OrderController) GetAllOrders(ctx *gin.Context) {
	response := o.orderService.GetAllOrders()

	params.WriteJsonResponse(ctx.Writer, response)
	// params.WriteJsonResponse(rw, response)
}

func (o *OrderController) UpdateOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")

	id, _ := strconv.Atoi(orderID)

	var req params.OrderUpdate

	// err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	err := ctx.BindJSON(&req)
	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, &response)
		return
	}

	response := o.orderService.UpdateOrder(&req, uint(id))
	log.Default().Printf("Request body of update order for id %d : %v", uint(id), &req)
	params.WriteJsonResponse(ctx.Writer, response)
}

func (o *OrderController) DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")

	id, _ := strconv.Atoi(orderID)
	response := o.orderService.DeleteOrder(uint(id))

	params.WriteJsonResponse(ctx.Writer, response)
}
