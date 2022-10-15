package repositories

import (
	"github.com/quangtran88/anifni-gateway/constant"
	"github.com/quangtran88/anifni-gateway/core/domain"
	"github.com/quangtran88/anifni-gateway/utils"
	userGRPC "github.com/quangtran88/anifni-grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type UserRepository struct {
	grpcClient userGRPC.UserServiceClient
}

func NewUserRepository() *UserRepository {

	uri := utils.GetEnv(constant.UserServiceEnvKey)
	log.Printf("connecting user service at: %s", uri)
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	client := userGRPC.NewUserServiceClient(conn)
	return &UserRepository{grpcClient: client}
}

func (repo UserRepository) FindById(id domain.ID) (domain.User, error) {
	return domain.User{Id: id}, nil
}

func (repo UserRepository) Ping() (string, error) {
	ctx, cancel := utils.InitGRPCContext()
	defer cancel()

	r, err := repo.grpcClient.Ping(ctx, &userGRPC.PingMessage{})
	if err != nil {
		log.Printf("could not ping: %v", err)
		return "", err
	}

	return r.GetMessage(), nil
}
