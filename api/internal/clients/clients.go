package clients

import (
	"github.com/MaksKazantsev/Chatter/api/internal/config"
	filesPkg "github.com/MaksKazantsev/Chatter/files/pkg/grpc"
	messagesPkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
	postsPkg "github.com/MaksKazantsev/Chatter/posts/pkg/grpc"
	userPkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	UserClient     UserClient
	MessagesClient MessagesClient
	FilesClient    FilesClient
	PostsClient    PostsClient
}

func Connect(cfg config.Services) Clients {
	var cli Clients

	// Dialing connection
	authCC, err := dial(cfg.AuthAddr)
	messagesCC, err := dial(cfg.MessagesAddr)
	filesCC, err := dial(cfg.FilesAddr)
	postsCC, err := dial(cfg.PostsAddr)
	if err != nil {
		panic("failed to dial: " + err.Error())
	}

	// Initializing clients
	cli.PostsClient = NewPosts(postsPkg.NewPostsClient(postsCC))
	cli.FilesClient = NewFiles(filesPkg.NewFilesClient(filesCC))
	cli.UserClient = NewUserAuth(userPkg.NewUserClient(authCC))
	cli.MessagesClient = NewMessages(messagesPkg.NewMessagesClient(messagesCC))

	return cli
}

func dial(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
