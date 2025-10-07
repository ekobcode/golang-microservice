package repository

import (
	"database/sql"
	"golang-microservice/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindAll() ([]entity.User, error)
	FindByID(id int64) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	return r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var u entity.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) FindByID(id int64) (*entity.User, error) {
	var u entity.User
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).
		Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) Update(user *entity.User) error {
	_, err := r.db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3",
		user.Name, user.Email, user.ID)
	return err
}

func (r *userRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
