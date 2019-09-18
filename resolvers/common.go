package resolvers

import (
	"github.com/Krystian19/cactus-bff/gql"
)

// Resolver : Global resolver
type Resolver struct{}
type queryResolver struct{ *Resolver }

// Query : global query resolver function
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}