package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/joho/godotenv"
	"github.com/lubie-koty/rpc-compute-service-simple/internal/config"
	"github.com/lubie-koty/rpc-compute-service-simple/internal/validate"
)

func main() {
	validate.InitValidate()
	envVars, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}
	err = config.InitConfig(envVars)
	if err != nil {
		fmt.Println("Error initializing config:", err)
		os.Exit(1)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	err = runApplication(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type application struct {
	logger *slog.Logger
	wg     sync.WaitGroup
}

func runApplication(logger *slog.Logger) error {
	app := &application{
		logger: logger,
	}

	return app.serveHTTP()
}
