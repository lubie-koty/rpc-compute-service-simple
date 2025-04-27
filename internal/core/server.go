package core

import (
	"log/slog"
	"os"
)

type Server struct {
	Host         string
	Port         int
	Logger       *slog.Logger
	ShutdownChan chan os.Signal
}
