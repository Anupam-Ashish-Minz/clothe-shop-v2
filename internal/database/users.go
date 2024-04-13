package database

import "fmt"

type User struct {
	ID       int64
	Name     string
	Email    string
	Password []byte
}

func (s *service) GetUserByEmail(email string) (User, error) {
	var user User
	row := s.db.QueryRow(`select id, name, email, password from "User" where email = $1`, email)
	row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if user.ID == 0 || user.Email == "" {
		return user, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *service) AddNewUser(user User) (int64, error) {
	if user.Name == "" || user.Email == "" || user.Password == nil || len(user.Password) == 0 {
		return 0, fmt.Errorf("missing fields in user object")
	}
	row := s.db.QueryRow(`insert into "User" (name, email, password) values ($1, $2, $3) returning id`,
		user.Name, user.Email, user.Password)
	var userID int64
	err := row.Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (s *service) GetUserById(userID int64) (User, error) {
	var user User
	row := s.db.QueryRow(`select id, name, email, password from "User" where id = $1`, userID)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	if user.ID == 0 || user.Email == "" {
		return user, fmt.Errorf("user not found")
	}
	return user, nil
}
