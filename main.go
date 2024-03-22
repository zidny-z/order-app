package main

import (
	"fmt"
	"order-app/config"
)

func main() {
	// print env
	fmt.Println(config.GetConfigDB())	// print db config
	config.GetConfigDB()
	// print port
	config.GetConfigPort()

}