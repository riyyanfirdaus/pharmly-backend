package main

import (
	"log"
	server "pharmly-backend/internal/Server"
)

func main() {
	application, err := server.NewApp()
	if err != nil {
		log.Fatal("Failed to initialize application:", err)
	}

	if err := application.Initialize(); err != nil {
		log.Fatal("Failed to initialize components:", err)
	}

	if err := application.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}

	defer func() {
		if err := application.Shutdown(); err != nil {
			log.Fatal("Failed to shutdown gracefully:", err)
		}
	}()
}
