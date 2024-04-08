package database

import "fmt"

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
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
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return 0, fmt.Errorf("missing fields in user object")
	}
	res, err := s.db.Exec(`insert into "User" (name, email, password) values ($1, $2, $3)`,
		user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
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
