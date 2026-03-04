package main

import (
	"context"
	"go-auth/internal/app"
	httpserver "go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
	// "github.com/gin-gonic/gin"
)

func main() {

	// this is a root context
	ctx := context.Background()

	a, err := app.New(ctx)

	if err != nil {
		log.Fatalf("StartUp failed: %v", err)
	}

	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Printf("App close failed: %v", err)
		}
	}()

	router := httpserver.NewRouter(a)

	// standard go type that runs a http server
	server := &http.Server{
		Addr:              ":5000",
		Handler:           router,
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
