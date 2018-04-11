package main

import (
	. "AppApi/project/database"
)

func main() {
	defer SqlDB.Close()
	router := initRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
