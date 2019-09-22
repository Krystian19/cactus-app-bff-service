package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-bff/gql"
	"github.com/Krystian19/cactus-core/proto"
)

type genreResolver struct{ *Resolver }

func (r *queryResolver) Genre(ctx context.Context, id *int) (*proto.Genre, error) {
	conn, client, err := genreServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	request := &proto.GenreRequest{}

	if id != nil {
		request.Id = int64(*id)
	}

	response, err := client.Genre(ctx, request)

	if err != nil {
		return nil, err
	}

	return response.Genre, nil
}

func (r *queryResolver) Genres(ctx context.Context, filter *gql.GenresFilter) (*gql.GenrePaginatedList, error) {
	conn, client, err := genreServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	request := &proto.GenresRequest{Query: &proto.GenreQuery{}}

	if filter != nil {
		if filter.Title != nil {
			request.Query.Title = *filter.Title
		}

		if filter.Limit != nil {
			request.Query.Limit = int64(*filter.Limit)
		}

		if filter.Offset != nil {
			request.Query.Offset = int64(*filter.Offset)
		}
	}

	response, err := client.Genres(ctx, request)

	if err != nil {
		return nil, err
	}

	return &gql.GenrePaginatedList{Rows: response.Genres, Count: int(response.Count)}, nil
}
