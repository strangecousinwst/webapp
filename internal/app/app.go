package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/strangecousinwst/webapp/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// TODO: Implement stores

	// TODO: Implement handlers

	app := &Application{
		Logger: logger,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status is available")
}
