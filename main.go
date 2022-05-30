package main

import (
	"go-restapi/controllers"
	"go-restapi/database"
	"go-restapi/repositories"
	"go-restapi/server"
	"go-restapi/services"
)

func main() {
	// port := ":8080"

	// log.Default().Println("Server running at port :", port)

	db := database.ConnectDB()
	// orderRepo := repositories.NewOrderRepo(db)
	orderRepo := repositories.NewOrderRepo(db)
	itemRepo := repositories.NewItemRepo(db)
	orderService := services.NewOrderService(orderRepo, itemRepo)
	orderController := controllers.NewOrderController(orderService)

	server := server.NewServer(orderController)
	server.StartServer()

}
