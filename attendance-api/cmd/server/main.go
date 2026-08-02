package main

import (
	"log"

	"attendance-api/internal/app"
)

func main() {

	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server starting on port %s", application.Config.Port)

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}

	defer application.Shutdown()
}
