package services

import (
	"github.com/quangtran88/anifni-gateway/core/domain"
	"github.com/quangtran88/anifni-gateway/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (srv UserService) Ping() (string, error) {
	return srv.repo.Ping()
}

func (srv UserService) Get(id domain.ID) (*domain.User, error) {
	return srv.repo.FindById(id)
}
