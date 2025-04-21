package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

// NewApplication initializes a new Application instance with a logger.
// Logger is to be used across the application instead
// of format print line.
func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Logger: logger,
	}

	return app, nil
}
