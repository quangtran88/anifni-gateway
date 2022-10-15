package ports

import "github.com/quangtran88/anifni-gateway/core/domain"

type UserRepository interface {
	Ping() (string, error)
	FindById(id domain.ID) (*domain.User, error)
}
