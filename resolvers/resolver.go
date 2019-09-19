package resolvers

import (
	"github.com/Krystian19/cactus-bff/gql"
)

// Resolver : Global resolver
type Resolver struct{}
type queryResolver struct{ *Resolver }

// Query : returns a QueryResolver struct which contains all query resolvers.
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

// Anime : returns a AnimeResolver struct which implements the gql.AnimeResolver interface.
func (r *Resolver) Anime() gql.AnimeResolver {
	return &animeResolver{r}
}
