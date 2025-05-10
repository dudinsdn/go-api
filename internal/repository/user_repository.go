package repository

import (
	"database/sql"
	"my/go-api/internal/model"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	Insert(user model.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Insert(user model.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name) VALUES (?, ?)", user.ID, user.Name)
	return err
}
