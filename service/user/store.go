package user

import (
	"database/sql"
	"fmt"
	"myapi/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {

	rows, err := s.db.Query("SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		u, err := ScanUserIntoRows(rows)
		if err != nil {
			return nil, err
		}
		return u, nil
	}
	return nil, fmt.Errorf("no user found")
}

func (s *Store) CreateUser(user types.User) error {
	if _, err := s.GetUserByEmail(user.Email); err == nil {
		return fmt.Errorf("user already exist")
	}

	_, err := s.db.Exec("INSERT INTO users (firstname,lastname,email,password) VALUES (?,?,?,?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		u, err := ScanUserIntoRows(rows)
		if err != nil {
			return nil, err
		}
		return u, nil
	}
	return nil, fmt.Errorf("no user found")
}

func ScanUserIntoRows(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)
	// Table columns: id, firstname, lastname, email, password, created_at
	err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}
