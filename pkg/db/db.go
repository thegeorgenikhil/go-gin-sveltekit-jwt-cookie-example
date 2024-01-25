package db

import (
	"database/sql"
)

type User struct {
	ID       int
	Email    string
	Password string
}

const createUserTable = `
CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL,
	password TEXT NOT NULL
);`

type UserDB struct {
	db *sql.DB
}

func New(file string) (*UserDB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(createUserTable); err != nil {
		return nil, err
	}

	return &UserDB{db: db}, nil
}

func (udb *UserDB) Close() error {
	return udb.db.Close()
}

// CreateUser inserts a new user into the database.
func (udb *UserDB) CreateUser(email string, password string) error {
	_, err := udb.db.Exec("INSERT INTO user (email, password) VALUES (?, ?)", email, password)
	return err
}

// GetUser retrieves a user by their email from the database.
func (udb *UserDB) GetUser(email string) (*User, error) {
	row := udb.db.QueryRow("SELECT id, email, password FROM user WHERE email = ?", email)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CheckIfUserExists checks if a user with the given email exists in the database.
func (udb *UserDB) CheckIfUserExists(email string) (bool, error) {
	var exists bool
	err := udb.db.QueryRow("SELECT EXISTS(SELECT 1 FROM user WHERE email = ?)", email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
