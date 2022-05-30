package server

import (
	"go-restapi/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	orderController *controllers.OrderController
}

func NewServer(order *controllers.OrderController) *Server {
	return &Server{
		orderController: order,
	}
}

func (s *Server) StartServer() {
	port := ":8080"

	router := gin.Default()

	router.GET("/orders", s.orderController.GetAllOrders)
	router.POST("/orders", s.orderController.CreateOrder)
	router.PUT("/orders/:orderID", s.orderController.UpdateOrder)
	router.DELETE("/orders/:orderID", s.orderController.DeleteOrder)

	log.Println("Server running at port", port)

	router.Run(port)

}
