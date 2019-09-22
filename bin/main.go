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

	gqlConfig := gql.Config{Resolvers: &resolvers.Resolver{}}

	http.Handle("/", handler.Playground("GraphQL playground", "/playground_graphql"))
	http.Handle("/playground_graphql", handler.GraphQL(gql.NewExecutableSchema(gqlConfig)))

	// Limit the query complexity of the endpoint exposed to the outside
	http.Handle("/graphql", handler.GraphQL(
		gql.NewExecutableSchema(gqlConfig),
		handler.ComplexityLimit(5), // GQL query complexity limit
	))

	log.Printf("GraphQL playground @ http://localhost:%s/", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

// EnvVarsCheck : Checks that important ENV vars are set
func EnvVarsCheck() {
	if os.Getenv("CACTUS_CORE_URL") == "" {
		panic("CACTUS_CORE_URL env var is not set")
	}
}
