package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"clothe-shop-v2/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port   int
	db     database.Service
	secret string
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	secret := os.Getenv("SECRET")
	NewServer := &Server{
		port:   port,
		db:     database.New(),
		secret: secret,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
