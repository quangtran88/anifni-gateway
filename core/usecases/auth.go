package usecases

import (
	"context"
	"encoding/json"
	"github.com/quangtran88/anifni-base/libs/event"
	"github.com/quangtran88/anifni-gateway/core/ports"
	"gopkg.in/errgo.v2/errors"
)

type AuthUseCase struct {
	userSrv ports.UserService
	event   ports.EventProducer
}

func NewAuthUseCase(userSrv ports.UserService, event ports.EventProducer) *AuthUseCase {
	return &AuthUseCase{userSrv, event}
}

func (uc AuthUseCase) PreRegisterUser(ctx context.Context, in ports.PreRegisterUserInput) (bool, error) {
	ok, err := uc.userSrv.CheckDuplicated(ctx, in.Email)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New("User with this email is existed")
	}

	mess, _ := json.Marshal(event.SendOTPRequestMessage{Email: in.Email})
	err = uc.event.Produce(ctx, event.SendOTPRequestTopic, in.Email, string(mess))
	if err != nil {
		return false, err
	}

	return true, nil
}
