package user

import (
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

// User represents a user with a balance
type User struct {
	ID      string `sql:"id"`
	Name    string  `json:"name" sql:"name"`
	Balance float64 `json:"balance,omitempty" sql:"balance,omitempty"`
}	

const (
	sqlCreateNewUser = `
		INSERT INTO	users (id, name, balance) VALUES ($1, $2, $3)`

	sqlGetUsers = `
	SELECT id, name
	FROM users
	`

	sqlGetUserBalance = `
	SELECT balance
	FROM users
	WHERE ID = $1
	`
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

// CreateUser creates a new user in the database from the given API endpoint
func(s *Service) CreateUser(u *User) (*User, error) {
	uuid := uuid.New()

	_, err := s.db.Exec(sqlCreateNewUser, uuid, u.Name, u.Balance)
	if err != nil {
		return nil, err
	}

	u.ID = uuid.String()

	return u, nil
}

// GetUsers returns all users with their ID and Name, but no balance
func (s *Service) GetUsers() (*[]User, error) {
	var users []User
	err := s.db.Select(&users, sqlGetUsers)
	if err != nil{
		return nil, err
	}

	return &users, nil
}

//GetUserBalance gets the balance of a user
func (s *Service) GetUserBalance(userID string) (string, error) {
	var balance string

	err := s.db.Get(&balance, sqlGetUserBalance, userID)
	if err != nil {
		return "", err
	}

	return balance, nil
}