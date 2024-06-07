package clients

import (
	pkg "github.com/MaksKazantsev/SSO/auth/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	UserClient UserAuth
}

func Connect(addr string) Clients {
	var cli Clients
	authCC, err := dial(addr)
	if err != nil {
		panic("")
	}
	cli.UserClient = NewUserAuth(pkg.NewUserClient(authCC))
	return cli
}

func dial(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
