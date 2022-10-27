package store

import (
	"Forum/domain/schema"
	"database/sql"
)

var (
	CreateError = "error in creating new user"
	ReadError   = "error in finding user in the database"
	listError   = "error in getting candidate from the database"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (schema.User, error) {
	var user schema.User

	row, err := s.db.Query("SELECT *FROM Users WHERE email=" + email)
	if err != nil {
		return user, err
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&user.UserName, &user.UserEmail, &user.Passwd)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func (s *Store) CreateUser(u schema.User) error {
	stmt, err := s.db.Prepare("INSERT INTO Users(username, email, passwd) VALUES (?,?,?)")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(u.UserName, u.UserEmail, u.Passwd); err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteUser(u schema.User) error {
	stmt, err := s.db.Prepare("DELETE FROM Users where username=?")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(u.UserName); err != nil {
		return err
	}

	return nil
}

func (s *Store) ListUsers() ([]schema.User, error) {

	var users []schema.User

	rows, err := s.db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var username, email, passwd string
		err = rows.Scan(&username, email, &passwd)
		if err != nil {
			return nil, err
		}

		users = append(users, schema.User{UserName: username,
			UserEmail: email, Passwd: passwd})
	}

	return users, nil
}
