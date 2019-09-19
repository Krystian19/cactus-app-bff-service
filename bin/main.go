package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/Krystian19/cactus-bff/gql"
	"github.com/Krystian19/cactus-bff/resolvers"
)

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "3000"
	}

	// Check for important env vars
	EnvVarsCheck()

	gqlHandler := handler.GraphQL(
		gql.NewExecutableSchema(gql.Config{Resolvers: &resolvers.Resolver{}}),
		handler.ComplexityLimit(5), // GQL query complexity limit
	)

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", gqlHandler)

	log.Printf("GraphQL playground @ http://localhost:%s/", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

// EnvVarsCheck : Checks that important ENV vars are set
func EnvVarsCheck() {
	CactusCoreURL := os.Getenv("CACTUS_CORE_URL")

	if CactusCoreURL == "" {
		panic("CACTUS_CORE_URL env var is not set")
	}
}
