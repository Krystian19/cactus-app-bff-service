package resolvers

import (
	"github.com/Krystian19/cactus-core/proto"
)

type releaseResolver struct{ *Resolver }

func releaseServiceClient() (client proto.ReleaseServiceClient, err error) {
	conn, err := InitGRPCConnection()

	if err != nil {
		return nil, err
	}

	return proto.NewReleaseServiceClient(conn), nil
}
