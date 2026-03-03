package main

import (
	httpserver "go-auth/internal"
	"log"
	"net/http"
	"time"

	// "github.com/gin-gonic/gin"
)

func main() {
	router := httpserver.NewRouter()

	// standard go type that runs a http server
	server := &http.Server{
		Addr: ":5000",
		Handler: router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("API running on %s ", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("Server closed: %v", err)
			return
		}
		log.Fatalf("Server error: %v", err)
	}
	
}