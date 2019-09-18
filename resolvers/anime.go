package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-core/proto"
)

// AnimeServiceClient : Returns a client for the ChartService
func AnimeServiceClient() (client proto.AnimeServiceClient, err error) {
	conn, err := InitGRPCConnection()

	if err != nil {
		return nil, err
	}

	return proto.NewAnimeServiceClient(conn), nil
}

func (r *queryResolver) Anime(ctx context.Context) (*proto.Anime, error) {
	client, err := AnimeServiceClient()

	if err != nil {
		return nil, err
	}

	response, err := client.Anime(ctx, &proto.Empty{})

	if err != nil {
		return nil, err
	}

	return response.Anime, nil
}
