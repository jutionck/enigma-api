package main

import (
	"github.com/jutionck/enigma-bank-api/config"
	"github.com/jutionck/enigma-bank-api/delivery"
)

func main() {

	db, _ := config.ConfigDb()
	router := config.CreateRouter()
	delivery.Init(router, db)
	config.RunServer(router)
}
