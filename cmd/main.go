package main

import (
	"log/slog"
	"os"
)

func main() {

	config := config{
		adress: ":8080",
		db:     dbConfig{},
	}

	api := application{
		config: config,
	}

	slog := slog.New(slog.NewTextHandler(os.Stdout, nil))

	api.run(api.mount())

	if err := api.run(api.mount()); err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}

}
