package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	if user.Password != password {
		c.String(http.StatusUnauthorized, "incorrect email or password")
		return
	}
	tokenString, err := createToken(user.ID, s.secret)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to authenitcate user")
		return
	}
	c.SetCookie("auth-token", tokenString, 86400, "/", "localhost", false, true)
}
