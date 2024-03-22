package main

import (
	"fmt"
	"order-app/controller"
	"order-app/database"
	router "order-app/routers"
)

func main() {
	db, err := database.Start()
	if err != nil {
		fmt.Println("Error starting database")
		return
	}

	cont := controller.New(db)

	err = router.StartServer(cont)
	if err != nil {
		fmt.Println("Error starting server")
		return
	}
}