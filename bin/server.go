package main

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Krystian19/cactus-bff/gql"
	"github.com/Krystian19/cactus-bff/resolvers"
)

const (
	gqlPlaygroundURL = "/play_graphql"
	gqlEndPoint      = "/q"
)

func newDefaultServer() *handler.Server {
	hldr := handler.New(gql.NewExecutableSchema(gql.Config{Resolvers: &resolvers.Resolver{}}))

	hldr.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})

	hldr.AddTransport(transport.Options{})
	hldr.AddTransport(transport.GET{})
	hldr.AddTransport(transport.POST{})
	hldr.AddTransport(transport.MultipartForm{})

	hldr.SetQueryCache(lru.New(1000))

	hldr.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return hldr
}

// Server : Runs a new GraphQL server
func Server(port string) error {
	// Graphql endpoint for internal use
	playSrv := newDefaultServer()
	playSrv.Use(apollotracing.Tracer{})
	playSrv.Use(extension.Introspection{})

	http.Handle("/", playground.Handler("Playground", gqlPlaygroundURL))
	http.Handle(gqlPlaygroundURL, playSrv)

	// Graphql endpoint exposed to the outside world
	outsideSrv := newDefaultServer()
	http.Handle(gqlEndPoint, outsideSrv)

	return http.ListenAndServe(":"+port, nil)
}
