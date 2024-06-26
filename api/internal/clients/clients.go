package clients

import (
	"github.com/MaksKazantsev/Chatter/api/internal/config"
	filesPkg "github.com/MaksKazantsev/Chatter/files/pkg/grpc"
	messagesPkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
	userPkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	UserClient     User
	MessagesClient Messages
	FilesClient    Files
}

func Connect(cfg config.Services) Clients {
	var cli Clients
	authCC, err := dial(cfg.AuthAddr)
	if err != nil {
		panic("failed to dial to auth")
	}
	messagesCC, err := dial(cfg.MessagesAddr)
	if err != nil {
		panic("failed to dial to messages")
	}
	filesCC, err := dial(cfg.FilesAddr)
	cli.FilesClient = NewFiles(filesPkg.NewFilesClient(filesCC))
	cli.UserClient = NewUserAuth(userPkg.NewUserClient(authCC))
	cli.MessagesClient = NewMessages(messagesPkg.NewMessagesClient(messagesCC))
	return cli
}

func dial(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
