package services

import (
	"go-restapi/models"
	"go-restapi/params"
	"go-restapi/repositories"
	"net/http"
)

type OrderServices struct {
	OrderRepo repositories.OrderRepo
}

func NewOrderService(OrderRepo repositories.OrderRepo) *OrderServices {
	return &OrderServices{
		OrderRepo: OrderRepo,
	}
}

func (o *OrderServices) GetAllOrders() *params.Response {
	orders, err := o.OrderRepo.GetAllOrders()
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: orders,
	}
}

func (o *OrderServices) CreateOrder(req *params.OrderCreate) *params.Response {
	order := &models.Order{
		CustomerName: req.CustomerName,
	}
	err := o.OrderRepo.CreateOrder(order)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "CREATED SUCCESS",
	}
}

func (o *OrderServices) UpdateOrder(req *params.OrderUpdate, id uint) *params.Response {
	order := &models.Order{
		CustomerName: req.CustomerName,
	}
	// id := req.ID
	err := o.OrderRepo.UpdateOrder(order, id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "UPDATED SUCCESS",
	}
}

func (o *OrderServices) DeleteOrder(id uint) *params.Response {
	err := o.OrderRepo.DeleteOrder(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "DELETED SUCCESS",
	}
}
