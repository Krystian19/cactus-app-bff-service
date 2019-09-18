package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-bff/gql"
)

func (r *queryResolver) Anime(ctx context.Context) (*gql.Anime, error) {
	return &gql.Anime{ID: "1", Title: "Here we are "}, nil
}
