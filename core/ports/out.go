package ports

import (
	"context"
	"github.com/quangtran88/anifni-gateway/core/domain"
)

type UserService interface {
	Ping(ctx context.Context) (string, error)
	CheckDuplicated(ctx context.Context, email string) (bool, error)
	Create(ctx context.Context, in CreateUserInput) (*domain.User, error)
}

type EnvManager interface {
	GetEnv(key string) string
	GetEnvDefault(key string, defaultValue string) string
}

type AuthService interface {
	CheckEmailOTP(ctx context.Context, code string, email string) (bool, error)
	SendEmailOTP(ctx context.Context, email string) (bool, error)
}

type TokenService interface {
	GetPID(ctx context.Context, domain string, prefix string) (string, error)
}
