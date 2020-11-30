package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-bff/gql"
	"github.com/Krystian19/cactus-core/proto"
)

type episodeSubtitleResolver struct{ *Resolver }

// EpisodeSubtitle : returns a EpisodeSubtitle struct which implements the gql.EpisodeSubtitleResolver interface.
func (r *Resolver) EpisodeSubtitle() gql.EpisodeSubtitleResolver {
	return &episodeSubtitleResolver{r}
}

func (r *episodeSubtitleResolver) Language(ctx context.Context, parent *proto.EpisodeSubtitle) (*proto.Language, error) {
	conn, client, err := languageServiceClient()
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	response, err := client.Language(ctx, &proto.LanguageRequest{Id: parent.LanguageId})

	if err != nil {
		return nil, err
	}

	return response.Language, nil
}
