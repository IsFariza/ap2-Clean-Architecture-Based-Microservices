package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	delivery "github.com/IsFariza/doctor-service/internal/delivery/http"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	uri := os.Getenv("MONGO_DB")
	port := os.Getenv("PORT")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	router := delivery.NewRouter(client)
	log.Printf("Server starting on port %s", port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	disconnectCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Disconnect(disconnectCtx); err != nil {
		log.Fatalf("MongoDB Disconnect Error: %v", err)
	}

	log.Println("Database connection closed.")
}
