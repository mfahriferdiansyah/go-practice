package main

import (
	"go-gin-gorm-mysql/database"
	router "go-gin-gorm-mysql/routers"
)

func main() {
	database.ConnectDatabase()
	r := router.RouterSetup()
	r.Run("127.0.0.1:9090")
}
