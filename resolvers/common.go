package resolvers

import (
	"os"

	"github.com/Krystian19/cactus-core/proto"
	"google.golang.org/grpc"
)

// InitGRPCConnection : Returns a grpc connection to the core service
func InitGRPCConnection() (conn *grpc.ClientConn, err error) {
	return grpc.Dial(os.Getenv("CACTUS_CORE_URL"), grpc.WithInsecure())
}

func animeServiceClient() (*grpc.ClientConn, proto.AnimeServiceClient, error) {
	conn, err := InitGRPCConnection()

	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, proto.NewAnimeServiceClient(conn), nil
}
