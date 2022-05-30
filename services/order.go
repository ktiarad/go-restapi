package services

import (
	"go-restapi/models"
	"go-restapi/params"
	"go-restapi/repositories"
	"log"
	"net/http"
)

type OrderServices struct {
	OrderRepo repositories.OrderRepo
	ItemRepo  repositories.ItemRepo
}

func NewOrderService(OrderRepo repositories.OrderRepo, ItemRepo repositories.ItemRepo) *OrderServices {
	return &OrderServices{
		OrderRepo: OrderRepo,
		ItemRepo:  ItemRepo,
	}
}

func (o *OrderServices) GetAllOrders() *params.Response {
	orders, err := o.OrderRepo.GetAllOrders()
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when get orders",
			AdditionalInfo: err.Error(),
		}
	}

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when get items",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: orders,
	}
}

func (o *OrderServices) CreateOrder(req *params.OrderCreate) *params.Response {

	// Get items value
	var items []models.Item
	for _, item := range req.Items {
		var currentItem models.Item
		currentItem.ItemCode = item.ItemCode
		currentItem.Description = item.Description
		currentItem.Quantity = item.Quantity
		// currentItem.OrderID =
		items = append(items, currentItem)
	}
	order := &models.Order{
		CustomerName: req.CustomerName,
		Items:        &items,
	}
	id, err := o.OrderRepo.CreateOrder(order)

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
		Payload: id,
	}
}

func (o *OrderServices) UpdateOrder(req *params.OrderUpdate, id uint) *params.Response {
	var items []models.Item
	for _, item := range req.Items {
		var currentItem models.Item

		currentItem.ID = item.ID
		currentItem.ItemCode = item.ItemCode
		currentItem.Description = item.Description
		currentItem.Quantity = item.Quantity

		items = append(items, currentItem)
		log.Default().Printf("current item :%v", currentItem)
		err := o.ItemRepo.UpdateItem(&currentItem, id)
		if err != nil {
			log.Default().Printf("error at UpdateItem :%v ", err)
		}
	}

	order := &models.Order{
		CustomerName: req.CustomerName,
		Items:        &items,
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
