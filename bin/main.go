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

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &resolvers.Resolver{}})))

	log.Printf("GraphQL playground @ http://localhost:%s/", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
