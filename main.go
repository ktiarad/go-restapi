package main

import (
	"log"
)

func main() {
	port := ":8080"

	log.Default().Println("Server running at port :", port)

	// db := database.ConnectDB()
	// orderRepo := repositories.NewOrderRepo(db)

}
