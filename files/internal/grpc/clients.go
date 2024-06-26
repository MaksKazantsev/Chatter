package grpc

import (
	"github.com/MaksKazantsev/Chatter/files/internal/config"
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	UserClient User
}

func Connect(cfg config.Services) Clients {
	var cli Clients
	authCC, err := dial(cfg.AuthAddr)
	if err != nil {
		panic("failed to dial")
	}
	cli.UserClient = NewUser(pkg.NewUserClient(authCC))
	return cli
}

func dial(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
