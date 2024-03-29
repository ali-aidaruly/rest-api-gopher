package store

import "github.com/Nurshat0092/rest-api-gopher/internal/app/model"

// UserRepository ..
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
