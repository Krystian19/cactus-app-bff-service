package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-bff/gql"
	"github.com/Krystian19/cactus-core/proto"
)

type episodeResolver struct{ *Resolver }

func episodeServiceClient() (client proto.EpisodeServiceClient, err error) {
	conn, err := InitGRPCConnection()

	if err != nil {
		return nil, err
	}

	return proto.NewEpisodeServiceClient(conn), nil
}

func (r *queryResolver) Episode(ctx context.Context, id *int) (*proto.Episode, error) {
	client, err := episodeServiceClient()

	if err != nil {
		return nil, err
	}

	request := &proto.EpisodeRequest{}

	if id != nil {
		request.Id = int64(*id)
	}

	response, err := client.Episode(ctx, request)

	if err != nil {
		return nil, err
	}

	return response.Episode, nil
}

func (r *queryResolver) Episodes(ctx context.Context, filter *gql.EpisodesFilter) (*gql.EpisodePaginatedList, error) {
	client, err := episodeServiceClient()

	if err != nil {
		return nil, err
	}

	request := &proto.EpisodesRequest{Query: &proto.EpisodeQuery{}}

	if filter != nil {
		if filter.ReleaseID != nil {
			request.Query.ReleaseId = int64(*filter.ReleaseID)
		}

		if filter.Limit != nil {
			request.Query.Limit = int64(*filter.Limit)
		}

		if filter.Offset != nil {
			request.Query.Offset = int64(*filter.Offset)
		}
	}

	response, err := client.Episodes(ctx, request)

	if err != nil {
		return nil, err
	}

	return &gql.EpisodePaginatedList{Rows: response.Episodes, Count: int(response.Count)}, nil
}

func (r *queryResolver) HottestEpisodes(ctx context.Context, limit *int, offset *int) (*gql.EpisodePaginatedList, error) {
	client, err := episodeServiceClient()

	if err != nil {
		return nil, err
	}

	request := &proto.EpisodesRequest{Query: &proto.EpisodeQuery{}}

	if limit != nil {
		request.Query.Limit = int64(*limit)
	}

	if offset != nil {
		request.Query.Offset = int64(*offset)
	}

	response, err := client.HottestEpisodes(ctx, request)

	if err != nil {
		return nil, err
	}

	return &gql.EpisodePaginatedList{Rows: response.Episodes, Count: int(response.Count)}, nil
}

func (r *queryResolver) NewestEpisodes(ctx context.Context, limit *int, offset *int) (*gql.EpisodePaginatedList, error) {
	client, err := episodeServiceClient()

	if err != nil {
		return nil, err
	}

	request := &proto.EpisodesRequest{Query: &proto.EpisodeQuery{}}

	if limit != nil {
		request.Query.Limit = int64(*limit)
	}

	if offset != nil {
		request.Query.Offset = int64(*offset)
	}

	response, err := client.NewestEpisodes(ctx, request)

	if err != nil {
		return nil, err
	}

	return &gql.EpisodePaginatedList{Rows: response.Episodes, Count: int(response.Count)}, nil
}

func (r *episodeResolver) Release(ctx context.Context, parent *proto.Episode) (*proto.Release, error) {
	client, err := releaseServiceClient()

	if err != nil {
		return nil, err
	}

	request := &proto.ReleaseRequest{Id: parent.ReleaseId}
	response, err := client.Release(ctx, request)

	if err != nil {
		return nil, err
	}

	return response.Release, nil
}
