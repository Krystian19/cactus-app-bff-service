package resolvers

import (
	"os"

	"github.com/Krystian19/cactus-bff/gql"
	"google.golang.org/grpc"
)

// Resolver : Global resolver
type Resolver struct{}
type queryResolver struct{ *Resolver }

// Query : global query resolver function
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

// InitGRPCConnection : Returns a grpc connection to the core service
func InitGRPCConnection() (conn *grpc.ClientConn, err error) {
	return grpc.Dial(os.Getenv("CACTUS_CORE_URL"), grpc.WithInsecure())
}
