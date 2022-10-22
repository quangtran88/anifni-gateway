package serviceAdapters

import (
	"context"
	"github.com/quangtran88/anifni-base/libs/constants"
	"github.com/quangtran88/anifni-gateway/core/ports"
	"github.com/quangtran88/anifni-grpc/authentication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const AuthServiceTimeoutSec = 8

type AuthenticationService struct {
	grpcClient authGRPC.OTPServiceClient
	env        ports.EnvManager
}

func NewAuthenticationService(env ports.EnvManager) *AuthenticationService {
	authSrvURI := env.GetEnv(baseConstants.AuthenticationURIEnvKey)

	log.Printf("connecting auth service at: %s", authSrvURI)
	conn, err := grpc.Dial(authSrvURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}

	authSrv := authGRPC.NewOTPServiceClient(conn)

	return &AuthenticationService{authSrv, env}
}

func (srv AuthenticationService) CheckEmailOTP(ctx context.Context, code string, email string) (bool, error) {
	ctx, cancel := srv.wrapCtx(ctx)
	defer cancel()

	res, err := srv.grpcClient.CheckEmailOTP(ctx, &authGRPC.CheckEmailOTPInput{
		Code:  code,
		Email: email,
	})
	if err != nil {
		return false, err
	}

	return res.Ok, nil
}

func (srv AuthenticationService) SendEmailOTP(ctx context.Context, email string) (bool, error) {
	ctx, cancel := srv.wrapCtx(ctx)
	defer cancel()

	result, err := srv.grpcClient.SendEmailOTP(ctx, &authGRPC.SendEmailOTPInput{Email: email})
	if err != nil {
		return false, err
	}

	return result.Ok, nil
}

func (srv AuthenticationService) wrapCtx(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, AuthServiceTimeoutSec*time.Second)
}
