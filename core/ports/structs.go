package ports

import "github.com/quangtran88/anifni-gateway/core/domain"

type CreateUserInput struct {
	Pid       domain.PID
	Email     string
	Password  string
	LastName  string
	FirstName string
}

type EventMessage struct {
	Key   string
	Value string
}
