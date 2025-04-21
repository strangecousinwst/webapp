package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/strangecousinwst/webapp/internal/app"
)

// Entry point
func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("Application is running")

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         ":42069",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status is available")
}
