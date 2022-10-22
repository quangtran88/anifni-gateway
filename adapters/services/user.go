package serviceAdapters

import (
	"context"
	baseConstants "github.com/quangtran88/anifni-base/libs/constants"
	"github.com/quangtran88/anifni-gateway/core/domain"
	"github.com/quangtran88/anifni-gateway/core/ports"
	userGRPC "github.com/quangtran88/anifni-grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const UserServiceTimeoutSec = 8

type UserService struct {
	grpcClient userGRPC.UserServiceClient
	env        ports.EnvManager
}

func NewUserService(env ports.EnvManager) *UserService {
	userSrvURI := env.GetEnv(baseConstants.UserURIEnvKey)

	log.Printf("connecting user service at: %s", userSrvURI)
	conn, err := grpc.Dial(userSrvURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}

	userClient := userGRPC.NewUserServiceClient(conn)
	return &UserService{userClient, env}
}

func (srv UserService) Ping(ctx context.Context) (string, error) {
	ctx, cancel := srv.wrapCtx(ctx)
	defer cancel()

	result, err := srv.grpcClient.Ping(ctx, &userGRPC.PingInput{})
	if err != nil {
		return "", err
	}
	return result.Message, nil
}

func (srv UserService) CheckDuplicated(ctx context.Context, email string) (bool, error) {
	ctx, cancel := srv.wrapCtx(ctx)
	defer cancel()

	result, err := srv.grpcClient.CheckDuplicated(ctx, &userGRPC.CheckDuplicatedUserInput{Email: email})
	if err != nil {
		return false, err
	}

	return result.Ok, nil
}

func (srv UserService) Create(ctx context.Context, in ports.CreateUserInput) (*domain.User, error) {
	ctx, cancel := srv.wrapCtx(ctx)
	defer cancel()

	result, err := srv.grpcClient.Create(ctx, &userGRPC.CreateUserInput{
		Pid:       string(in.Pid),
		Email:     in.Email,
		Password:  in.Password,
		LastName:  in.LastName,
		FirstName: in.FirstName,
	})
	if err != nil {
		return nil, err
	}

	return &domain.User{
		Pid:       domain.PID(result.Pid),
		Name:      result.Name,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
	}, nil
}

func (srv UserService) wrapCtx(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, UserServiceTimeoutSec*time.Second)
}
