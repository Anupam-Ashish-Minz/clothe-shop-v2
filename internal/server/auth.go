package server

import (
	"clothe-shop-v2/internal/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func createToken(userID int64, SECRET []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user_id": userID})
	return token.SignedString(SECRET)
}

func parseToken(tokenString string, SECRET []byte) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return SECRET, nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("failed to parse claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("failed to parse claims: float conversion")
	}
	return int64(userID), nil
}

func (s *Server) UserLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		c.String(http.StatusBadRequest, "missing fields")
		return
	}
	user, err := s.db.GetUserByEmail(email)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "user not found, please check if the email is correct and user is signup before continuing")
		return
	}
	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		c.String(http.StatusUnauthorized, "incorrect email or password")
		return
	}
	tokenString, err := createToken(user.ID, s.secret)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to authenitcate user")
		return
	}
	c.SetCookie("auth-token", tokenString, 86400, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/products")
}

func (s *Server) UserSignup(c *gin.Context) {
	var user database.User
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	password := c.PostForm("password")
	if password == "" {
		c.String(http.StatusBadRequest, "password is required")
		return
	}
	var err error
	user.Password, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "invalid password")
	}
	if user.Name == "" || user.Email == "" || user.Password == nil || len(user.Password) == 0 {
		c.String(http.StatusBadRequest, "missing fields")
		return
	}
	userID, err := s.db.AddNewUser(user)
	if err != nil {
		log.Println(err)
		c.String(http.StatusConflict, "email already used")
		return
	}
	tokenString, err := createToken(userID, s.secret)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to authenticate user, try loggin")
		return
	}
	c.SetCookie("auth-token", tokenString, 86400, "/", "localhost", false, true)
}

func (s *Server) Authenticate(c *gin.Context) (int64, error) {
	tokenString, err := c.Cookie("auth-token")
	if err != nil {
		return 0, err
	}
	if tokenString == "" {
		return 0, fmt.Errorf("auth token not found")
	}
	userID, err := parseToken(tokenString, s.secret)
	if err != nil {
		return 0, err
	}

	user, err := s.db.GetUserById(userID)
	if err != nil {
		return 0, err
	}
	if user.ID == 0 || user.Email == "" {
		return 0, fmt.Errorf("user authentication failed")
	}
	return userID, nil
}
