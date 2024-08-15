package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hftamayo/gorptmservice/internal/handlers"
)

func main() {
	router := gin.Default()

	router.POST("/generate-report", handlers.GenerateReportHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8004"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Create a channel to listen for termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	log.Printf("Server is running on port %s", port)

	// Block until a termination signal is received
	<-quit
	log.Println("Shutting down server...")

	// Create a context with a timeout of 5 seconds
	// This will allow the server to gracefully shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
