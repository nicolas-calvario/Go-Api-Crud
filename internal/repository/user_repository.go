package repository

import (
	"database/sql"
	"errors"

	"github.com/nicolas-calvario/Go-Api-Crud/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int64) (*domain.User, error)
	GetAll() ([]domain.User, error)
	Update(user *domain.User) error
	Delete(id int64) error
}

type userRepository struct{
	db *sql.DB
}

func NewUserRepository(	db *sql.DB) UserRepository{
	return &userRepository{db}
}


func (r *userRepository) Create(user *domain.User) error  {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	return err
}

func (r *userRepository) GetByID(id int64) (*domain.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]domain.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Update(user *domain.User) error {
	query := "UPDATE users SET name=$1, email=$2 WHERE id=$3"
	_, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

func (r *userRepository) Delete(id int64) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.db.Exec(query, id)
	return err
}