package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Bharat1056/students-api/internal/config"
)

func main() {
	// load config
	cfg := config.Mustload()
	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("env:", cfg.Env)
		fmt.Println("storage path:", cfg.StoragePath)
		fmt.Println("storage path:", cfg.StoragePath)
		fmt.Println("http server:", cfg.HTTPServer)
		fmt.Println("http server addr:", cfg.HTTPServer.Addr)

		w.Write([]byte("Hello From Golang"))
	})
	// setup server
	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	go func() {
		slog.Info("server started", slog.String("address:",cfg.Addr))
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

		<-done

		// server stop logic

		slog.Info("shutting down the server")

		ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("Failed to Shutdown Server", slog.String("error", err.Error()))
		}

		slog.Info("Server Shutdown successfully")
}
