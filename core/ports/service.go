package ports

import "github.com/quangtran88/anifni-gateway/core/domain"

type UserService interface {
	Ping() (string, error)
	Get(id domain.ID) (domain.User, error)
}
