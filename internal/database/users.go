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
	row := s.db.QueryRow(`select id, name, email, password from User where email = ?`, email)
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
	res, err := s.db.Exec(`insert into User (name, email, password) values (?, ?, ?)`,
		user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
