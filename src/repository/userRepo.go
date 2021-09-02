package repository

import "cloud.google.com/go/firestore"

type UserRepository interface {
	Get(id string) (*User, error)
	Create(user User) (*User, error)
	Update(user User) (*User, error)
}

type userRepository struct {
	db *firestore.Client
}

func NewUserRepository(db *firestore.Client) UserRepository {
	return &userRepository{db: db}
}

func (ref *userRepository)Get(id string) (*User, error) {
	return nil, nil
}

func (ref *userRepository)Create(user User) (*User, error) {
	return  nil, nil
}

func (ref *userRepository)Update(user User) (*User, error) {
	return nil, nil
}
