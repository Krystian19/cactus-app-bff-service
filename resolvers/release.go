package resolvers

import (
	"github.com/Krystian19/cactus-core/proto"
)

type releaseResolver struct{ *Resolver }

func releaseServiceClient() (client proto.ReleaseServiceClient, err error) {
	conn, err := InitGRPCConnection()

	if err != nil {
		return nil, err
	}

	return proto.NewReleaseServiceClient(conn), nil
}

// func (r *queryResolver) Anime(ctx context.Context) (*proto.Anime, error) {
// 	client, err := releaseServiceClient()

// 	if err != nil {
// 		return nil, err
// 	}

// 	response, err := client.Anime(ctx, &proto.Empty{})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return response.Anime, nil
// }

// func (r *animeResolver) Releases(ctx context.Context, parent *proto.Anime) ([]*proto.Release, error) {
// 	return []*proto.Release{
// 		&proto.Release{ Id: 1, Title: "This is a testing release" },
// 	}, nil
// }
