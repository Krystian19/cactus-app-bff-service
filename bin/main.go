package main

import (
	"log"
	"os"
)

const (
	// Default port for the GraphQL Server
	defaultPort = "3000"
)

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = defaultPort
	}

	log.Printf("GraphQL server running @ http://localhost:%s/", PORT)
	if err := Server(PORT); err != nil {
		log.Fatal(err)
	}
}
