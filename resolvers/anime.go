package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-core/proto"
)

type animeResolver struct{ *Resolver }

func (r *queryResolver) Anime(ctx context.Context, id *int) (*proto.Anime, error) {
	conn, client, err := animeServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	request := &proto.AnimeRequest{}

	if id != nil {
		request.Id = int64(*id)
	}

	response, err := client.Anime(ctx, request)

	if err != nil {
		return nil, err
	}

	return response.Anime, nil
}
