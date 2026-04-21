package main

import (
	"log"

	"github.com/phuongtran6575/bookstore-be.git/internal/data"
	"github.com/phuongtran6575/bookstore-be.git/internal/di"
)

func main() {
	db := data.InitDB()

	engine := di.SetupDI(db)

	log.Println("Server started on :8080")
	engine.Run(":8080")
}
