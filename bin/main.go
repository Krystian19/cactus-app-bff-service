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

	// Check for important env vars
	EnvVarsCheck()

	log.Printf("GraphQL server running @ http://localhost:%s/", PORT)
	if err := Server(PORT); err != nil {
		log.Fatal(err)
	}
}

// EnvVarsCheck : Checks that important ENV vars are set
func EnvVarsCheck() {
	if os.Getenv("CACTUS_CORE_URL") == "" {
		panic("CACTUS_CORE_URL env var is not set")
	}
}
