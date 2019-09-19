package resolvers

import (
	"os"

	"google.golang.org/grpc"
)

// InitGRPCConnection : Returns a grpc connection to the core service
func InitGRPCConnection() (conn *grpc.ClientConn, err error) {
	return grpc.Dial(os.Getenv("CACTUS_CORE_URL"), grpc.WithInsecure())
}
