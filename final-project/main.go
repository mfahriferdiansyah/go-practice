package main

import (
	db "final-project/database"
	"final-project/router"
)

func main() {
	db.ConnectDatabase()
	r := router.RouterSetup()
	r.Run("127.0.0.1:9090")
}
