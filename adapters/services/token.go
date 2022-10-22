package serviceAdapters

import (
	"context"
	baseConstants "github.com/quangtran88/anifni-base/libs/constants"
	"github.com/quangtran88/anifni-gateway/core/ports"
	"github.com/quangtran88/anifni-grpc/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type TokenService struct {
	grpcClient tokenGRPC.TokenServiceClient
	env        ports.EnvManager
}

func NewTokenService(env ports.EnvManager) *TokenService {
	authSrvURI := env.GetEnv(baseConstants.TokenURIEnvKey)

	log.Printf("connecting token service at: %s", authSrvURI)
	conn, err := grpc.Dial(authSrvURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}

	tokenClient := tokenGRPC.NewTokenServiceClient(conn)
	return &TokenService{tokenClient, env}
}

func (srv TokenService) GetPID(ctx context.Context, domain string, prefix string) (string, error) {
	ctx, cancel := srv.wrapCtx(ctx)
	defer cancel()

	result, err := srv.grpcClient.GetPID(ctx, &tokenGRPC.GetPIDInput{Domain: domain, Prefix: prefix})
	if err != nil {
		return "", err
	}

	return result.Pid, nil
}

func (srv TokenService) wrapCtx(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, AuthServiceTimeoutSec*time.Second)
}
