package resolvers

import (
	"context"
	"log"

	"github.com/Krystian19/cactus-core/proto"
)

// AnimeServiceClient : Returns a client for the ChartService
func AnimeServiceClient() (client proto.AnimeServiceClient, err error) {
	conn, err := InitGRPCConnection()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return proto.NewAnimeServiceClient(conn), nil
}

func (r *queryResolver) Anime(ctx context.Context) (*proto.Anime, error) {
	return &proto.Anime{Id: int64(1), Title: "Here we are "}, nil
}
