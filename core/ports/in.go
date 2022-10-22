package ports

import (
	"context"
	"github.com/quangtran88/anifni-gateway/core/domain"
)

type AuthUseCase interface {
	RegisterUser(ctx context.Context, in domain.RegisterUserInput) (bool, error)
}
