package userRepository

import (
	user "users/Domain/Model"
)

type UserRepository interface {
	GetUsers() ([]user.User, error)
	GetUser(id string) (*user.User, error)
	UpdateUser(user user.User) error
}
