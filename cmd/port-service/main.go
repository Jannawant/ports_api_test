package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"portsApi/internal/config"
	"portsApi/internal/repository/inmem"
	"portsApi/internal/services"
	"portsApi/internal/transport"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	cfg := config.Read()

	portsStoreRepo := inmem.NewPortStore()
	portService := services.NewPortService(portsStoreRepo)
	httpServer := transport.NewHttpServer(portService)

	router := mux.NewRouter()
	router.HandleFunc("/port", httpServer.GetPort).Methods("GET")
	router.HandleFunc("/count", httpServer.CountPorts).Methods("GET")
	router.HandleFunc("/ports", httpServer.UploadPorts).Methods("POST")

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}
	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", cfg.HTTPAddr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped
	log.Printf("Goodbye")
	return nil
}
