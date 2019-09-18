package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-core/proto"
)

func (r *queryResolver) Anime(ctx context.Context) (*proto.Anime, error) {
	return &proto.Anime{Id: int64(1), Title: "Here we are "}, nil
}
