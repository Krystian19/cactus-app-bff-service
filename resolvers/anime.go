package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-core/proto"
)

type animeResolver struct{ *Resolver }

func animeServiceClient() (client proto.AnimeServiceClient, err error) {
	conn, err := InitGRPCConnection()

	if err != nil {
		return nil, err
	}

	return proto.NewAnimeServiceClient(conn), nil
}

func (r *queryResolver) Anime(ctx context.Context) (*proto.Anime, error) {
	client, err := animeServiceClient()

	if err != nil {
		return nil, err
	}

	response, err := client.Anime(ctx, &proto.Empty{})

	if err != nil {
		return nil, err
	}

	return response.Anime, nil
}

func (r *animeResolver) Releases(ctx context.Context, parent *proto.Anime) ([]*proto.Release, error) {
	return []*proto.Release{
		&proto.Release{ Id: 1, Title: "This is a testing release" },
	}, nil
}
