package ports

import (
	"context"
)

type AuthUseCase interface {
	PreRegisterUser(ctx context.Context, in PreRegisterUserInput) (bool, error)
}
