package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-bff/gql"
	"github.com/Krystian19/cactus-core/proto"
)

type animeResolver struct{ *Resolver }

// Anime : returns a AnimeResolver struct which implements the gql.AnimeResolver interface.
func (r *Resolver) Anime() gql.AnimeResolver {
	return &animeResolver{r}
}

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

func (r *animeResolver) Releases(ctx context.Context, parent *proto.Anime) ([]*proto.Release, error) {
	conn, client, err := releaseServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	response, err := client.Releases(ctx, &proto.ReleasesRequest{Query: &proto.ReleaseQuery{AnimeId: parent.Id}})

	if err != nil {
		return nil, err
	}

	return response.Releases, nil
}
