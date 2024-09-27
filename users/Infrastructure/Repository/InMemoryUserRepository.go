package inMemoryUserRepository

import (
	"sync"
	user "users/Domain/Model"
)

var userRepository = make(map[string]user.User)

type UserRepository struct {
	users map[string]user.User
	lock  *sync.RWMutex
}

func loadFixtures() map[string]user.User {
	for _, user := range user.Users {
		userRepository[user.ID] = user // Populate the map
	}

	return userRepository
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: loadFixtures(),
		lock:  &sync.RWMutex{},
	}
}

func (m UserRepository) GetUser(id string) (*user.User, bool) {
	user, error := m.users[id]
	return &user, error
}

func (m UserRepository) GetUsers() map[string]user.User {
	return m.users
}

func (m UserRepository) UpdateUser(user user.User) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	m.users[user.ID] = user
}
