package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-bff/gql"
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

func (r *queryResolver) Release(ctx context.Context, id *int) (*proto.Release, error) {
	client, err := releaseServiceClient()

	if err != nil {
		return nil, err
	}

	request := &proto.ReleaseRequest{}

	if id != nil {
		request.Id = int64(*id)
	}

	response, err := client.Release(ctx, request)

	if err != nil {
		return nil, err
	}

	return response.Release, nil
}

func (r *queryResolver) Releases(ctx context.Context, filter *gql.ReleasesFilter) (*gql.ReleasePaginatedList, error) {
	client, err := releaseServiceClient()

	if err != nil {
		return nil, err
	}

	request := &proto.ReleasesRequest{Query: &proto.ReleaseQuery{}}

	if filter != nil {
		if filter.AnimeID != nil {
			request.Query.AnimeId = int64(*filter.AnimeID)
		}

		if filter.Limit != nil {
			request.Query.Limit = int64(*filter.Limit)
		}

		if filter.Offset != nil {
			request.Query.Offset = int64(*filter.Offset)
		}
	}

	response, err := client.Releases(ctx, request)

	if err != nil {
		return nil, err
	}

	return &gql.ReleasePaginatedList{Rows: response.Releases, Count: int(response.Count)}, nil
}
