package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-bff/gql"
	"github.com/Krystian19/cactus-core/proto"
)

type releaseResolver struct{ *Resolver }

func (r *queryResolver) Release(ctx context.Context, id *int) (*proto.Release, error) {
	conn, client, err := releaseServiceClient()
	defer conn.Close()

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
	conn, client, err := releaseServiceClient()
	defer conn.Close()

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

func (r *queryResolver) RandomRelease(ctx context.Context) (*proto.Release, error) {
	conn, client, err := releaseServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	response, err := client.RandomRelease(ctx, &proto.Empty{})

	if err != nil {
		return nil, err
	}

	return response.Release, nil
}

func (r *queryResolver) AiringReleases(ctx context.Context) ([]*proto.Release, error) {
	conn, client, err := releaseServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	response, err := client.AiringReleases(ctx, &proto.Empty{})

	if err != nil {
		return nil, err
	}

	return response.Releases, nil
}

func (r *releaseResolver) EpisodeCount(ctx context.Context, parent *proto.Release) (int, error) {
	conn, client, err := episodeServiceClient()
	defer conn.Close()

	if err != nil {
		return 0, err
	}

	response, err := client.EpisodeCount(ctx, &proto.EpisodeCountRequest{ReleaseId: parent.Id})

	if err != nil {
		return 0, err
	}

	return int(response.Count), nil
}

func (r *releaseResolver) LatestEpisode(ctx context.Context, parent *proto.Release) (*proto.Episode, error) {
	conn, client, err := episodeServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	request := &proto.LatestEpisodeRequest{ReleaseId: parent.Id}
	response, err := client.LatestEpisode(ctx, request)

	if err != nil {
		return nil, err
	}

	return response.Episode, nil
}

func (r *releaseResolver) Anime(ctx context.Context, parent *proto.Release) (*proto.Anime, error) {
	conn, client, err := animeServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	request := &proto.AnimeRequest{Id: parent.AnimeId}
	response, err := client.Anime(ctx, request)

	if err != nil {
		return nil, err
	}

	return response.Anime, nil
}
