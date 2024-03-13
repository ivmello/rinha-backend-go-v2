package core

import (
	"database/sql"
)

type UserRepository interface {
	Update(user User)
	Get(id int) User
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Update(user User) {
}

func (u *userRepository) Get(id int) User {
	return User{}
}
