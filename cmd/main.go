package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/vitconduck/fun/internal/adapter/handler/http"
	"github.com/vitconduck/fun/internal/adapter/postgres/repository"
	"github.com/vitconduck/fun/internal/core/service"
	"github.com/vitconduck/fun/pkg/configs"
	"github.com/vitconduck/fun/pkg/postgres"
)

func main() {
	cfg, err := configs.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	ctx := context.Background()

	db, err := postgres.New(ctx, cfg.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)

	router, err := http.NewRouter(cfg.HTTP, *userHandler)

	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", cfg.HTTP.URL, cfg.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}

}
