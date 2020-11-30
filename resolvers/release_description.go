package resolvers

import (
	"context"

	"github.com/Krystian19/cactus-bff/gql"
	"github.com/Krystian19/cactus-core/proto"
)

type releaseDescriptionResolver struct{ *Resolver }

// ReleaseDescription : returns a ReleaseDescriptionResolver struct which implements the gql.ReleaseDescriptionResolver interface.
func (r *Resolver) ReleaseDescription() gql.ReleaseDescriptionResolver {
	return &releaseDescriptionResolver{r}
}

func (r *releaseDescriptionResolver) Language(ctx context.Context, parent *proto.ReleaseDescription) (*proto.Language, error) {
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
