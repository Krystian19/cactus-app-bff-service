package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-bff/gql"
	"github.com/Krystian19/cactus-core/proto"
)

type episodeResolver struct{ *Resolver }

func (r *queryResolver) Episode(ctx context.Context, id *int) (*proto.Episode, error) {
	conn, client, err := episodeServiceClient()
	defer conn.Close()

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
	conn, client, err := episodeServiceClient()
	defer conn.Close()

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
	conn, client, err := episodeServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	request := &proto.PaginationRequest{}

	if limit != nil {
		request.Limit = int64(*limit)
	}

	if offset != nil {
		request.Offset = int64(*offset)
	}

	response, err := client.HottestEpisodes(ctx, request)

	if err != nil {
		return nil, err
	}

	return &gql.EpisodePaginatedList{Rows: response.Episodes, Count: int(response.Count)}, nil
}

func (r *queryResolver) NewestEpisodes(ctx context.Context, limit *int, offset *int) (*gql.EpisodePaginatedList, error) {
	conn, client, err := episodeServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	request := &proto.EpisodesRequest{
		OrderBy: &proto.OrderBy{
			Field:      "created_at",
			Descending: true,
		},
		Query: &proto.EpisodeQuery{},
	}

	if limit != nil {
		request.Query.Limit = int64(*limit)
	}

	if offset != nil {
		request.Query.Offset = int64(*offset)
	}

	response, err := client.Episodes(ctx, request)

	if err != nil {
		return nil, err
	}

	return &gql.EpisodePaginatedList{Rows: response.Episodes, Count: int(response.Count)}, nil
}

func (r *episodeResolver) EarlierEpisode(ctx context.Context, parent *proto.Episode) (*proto.Episode, error) {
	conn, client, err := episodeServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	response, err := client.Episode(ctx, &proto.EpisodeRequest{
		ReleaseId: parent.ReleaseId,
		LessThan: &proto.LessThan{
			Field: "episode_order",
			Value: parent.EpisodeOrder,
		},
		OrderBy: &proto.OrderBy{
			Field:      "episode_order",
			Descending: true,
		},
	})

	if err != nil {
		return nil, err
	}

	return response.Episode, nil
}

func (r *episodeResolver) LaterEpisode(ctx context.Context, parent *proto.Episode) (*proto.Episode, error) {
	conn, client, err := episodeServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	response, err := client.Episode(ctx, &proto.EpisodeRequest{
		ReleaseId: parent.ReleaseId,
		GreaterThan: &proto.GreaterThan{
			Field: "episode_order",
			Value: parent.EpisodeOrder,
		},
		OrderBy: &proto.OrderBy{
			Field:      "episode_order",
			Descending: false,
		},
	})

	if err != nil {
		return nil, err
	}

	return response.Episode, nil
}

func (r *episodeResolver) Release(ctx context.Context, parent *proto.Episode) (*proto.Release, error) {
	conn, client, err := releaseServiceClient()
	defer conn.Close()

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

func (r *episodeResolver) EpisodeSubtitles(ctx context.Context, parent *proto.Episode) ([]*proto.EpisodeSubtitle, error) {
	conn, client, err := episodeServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	response, err := client.EpisodeSubtitles(ctx, &proto.EpisodeSubtitlesRequest{EpisodeId: parent.Id})

	if err != nil {
		return nil, err
	}

	return response.EpisodeSubtitles, nil
}

func (r *mutationResolver) EpisodeSeen(ctx context.Context, EpisodeID int) (bool, error) {
	conn, client, err := episodeServiceClient()
	defer conn.Close()

	if err != nil {
		return false, err
	}

	if _, err := client.EpisodeSeen(ctx, &proto.EpisodeSeenRequest{EpisodeId: int64(EpisodeID)}); err != nil {
		return false, err
	}

	return true, nil
}
