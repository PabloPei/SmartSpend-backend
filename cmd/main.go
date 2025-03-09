package main

import (
	"log"

	"github.com/PabloPei/SmartSpend-backend/conf"
	"github.com/PabloPei/SmartSpend-backend/db"
	"github.com/PabloPei/SmartSpend-backend/internal/api"
)

func main() {

	// PSQL Connection //

	postgresCfg := conf.InitPostgresSqlConfig()

	log.Println("Starting PostgreSQL connection...")

	db, err := db.NewPostgresStorage(postgresCfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the database")

	// API Server //

	log.Println("Starting Api Server...")

	serverCfg := conf.InitApiServerConfig()
	server := api.NewAPIServer(serverCfg, db)
	server.Run()
}
