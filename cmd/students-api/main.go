package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Bharat1056/students-api/internal/config"
	student "github.com/Bharat1056/students-api/internal/http/handler/students"
)

func main() {
	// load config
	cfg := config.Mustload()
	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /api/students", student.New())
	// setup server
	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	// instead of buffered if we set an un-buffered channel then issue we might face is
	// we sent many signal to the channel then also server is not gonna shutdown because it is theoritically infinite signal can absorb

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// run the server listen inside a go routine
	go func() {
		slog.Info("server started", slog.String("address:",cfg.Addr))
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

		// as the go routine is non-blocking so we need to block the server from here
		// for this we can use channel or mutex
		<-done

		// server stop logic

		slog.Info("shutting down the server")

		// we can use server.Shutdown but it might not be shutdown sometimes
		// sometimes it is waiting for infinite time
		// so for this we use context - nothing just handle those things

		// create a timeline in that context
		ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel() // this is basically garbage collector call of that context

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("Failed to Shutdown Server", slog.String("error", err.Error()))
		}

		slog.Info("Server Shutdown successfully")
}


// go run md/students-api/main.go -config config/loal.yaml - to run the application
