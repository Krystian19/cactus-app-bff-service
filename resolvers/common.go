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

func episodeServiceClient() (*grpc.ClientConn, proto.EpisodeServiceClient, error) {
	conn, err := InitGRPCConnection()

	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, proto.NewEpisodeServiceClient(conn), nil
}

func genreServiceClient() (*grpc.ClientConn, proto.GenreServiceClient, error) {
	conn, err := InitGRPCConnection()

	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, proto.NewGenreServiceClient(conn), nil
}

func releaseServiceClient() (*grpc.ClientConn, proto.ReleaseServiceClient, error) {
	conn, err := InitGRPCConnection()

	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, proto.NewReleaseServiceClient(conn), nil
}
