package main

import (
	"fmt"
	"log"

	"notes-api/internal/config"
	"notes-api/internal/db"
	"notes-api/internal/server"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	// FIX: Capture the database instead of discarding it with _
	// The database is needed to pass to NewRouter
	client, database, err := db.Connect(cfg)

	if err != nil {
		log.Fatalf("db connection error: %v", err)
	}

	defer func() {
		if err := db.Disconnect(client); err != nil {
			log.Printf("db disconnect error: %v", err)
		}
	}()

	// FIX: Pass the database to NewRouter as required by its signature
	router := server.NewRouter(database)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
