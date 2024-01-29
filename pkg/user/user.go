package user

import (
	"database/sql"

	"github.com/google/uuid"
)

// User represents a user with a balance
type User struct {
	ID      string
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}	

const (
	sqlCreateNewUser = `
		INSERT INTO	users (id, name, balance) VALUES ($1, $2, $3)`
)


type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetUser(userID int) {

}

func(s *Service) CreateUser(u *User) (*User, error) {
	uuid := uuid.New()

	_, err := s.db.Exec(sqlCreateNewUser, uuid, u.Name, u.Balance)
	if err != nil {
		return nil, err
	}

	u.ID = uuid.String()

	return u, nil
}