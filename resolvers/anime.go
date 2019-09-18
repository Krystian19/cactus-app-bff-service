package resolvers

import (
	"context"

	gqlgen "github.com/Krystian19/cactus-bff/gql"
)

func (r *queryResolver) Anime(ctx context.Context) (*gqlgen.Anime, error) {
	return &gqlgen.Anime{ID: "1", Title: "Here we are "}, nil
}
