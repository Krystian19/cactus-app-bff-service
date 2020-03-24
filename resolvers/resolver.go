package resolvers

import (
	"github.com/Krystian19/cactus-bff/gql"
)

// Resolver : Global resolver
type Resolver struct{}
type queryResolver struct{ *Resolver }

// type mutationResolver struct{ *Resolver }

// Query : returns a QueryResolver struct which contains all query resolvers.
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

// Mutation : returns a MutationResolver struct which contains all mutation resolvers.
// func (r *Resolver) Mutation() gql.MutationResolver {
// 	return &mutationResolver{r}
// }

// // Anime : returns a AnimeResolver struct which implements the gql.AnimeResolver interface.
// func (r *Resolver) Anime() gql.AnimeResolver {
// 	return &animeResolver{r}
// }

// // Release : returns a ReleaseResolver struct which implements the gql.ReleaseResolver interface.
// func (r *Resolver) Release() gql.ReleaseResolver {
// 	return &releaseResolver{r}
// }

// // ReleaseDescription : returns a ReleaseDescriptionResolver struct which implements the gql.ReleaseDescriptionResolver interface.
// func (r *Resolver) ReleaseDescription() gql.ReleaseDescriptionResolver {
// 	return &releaseDescriptionResolver{r}
// }

// // Episode : returns a EpisodeResolver struct which implements the gql.EpisodeResolver interface.
// func (r *Resolver) Episode() gql.EpisodeResolver {
// 	return &episodeResolver{r}
// }

// // EpisodeSubtitle : returns a EpisodeSubtitle struct which implements the gql.EpisodeSubtitleResolver interface.
// func (r *Resolver) EpisodeSubtitle() gql.EpisodeSubtitleResolver {
// 	return &episodeSubtitleResolver{r}
// }
