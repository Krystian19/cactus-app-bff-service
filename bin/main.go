package main

import "log"

const (
	// PORT : GraphQL Server port
	PORT = "3000"
)

func main() {
	log.Printf("GraphQL server running @ http://localhost:%s/", PORT)
	if err := Server(PORT); err != nil {
		log.Fatal(err)
	}
}
