package server

import (
	"fmt"
	"log"
	"os"

	"clothe-shop-v2/internal/database"

	"github.com/joho/godotenv"
)

type Product struct {
	Name        string
	Description string
	Price       int
	Gender      string
}

func setupTesting() (*Server, error) {
	godotenv.Load("../../.env")
	fmt.Println("starting setup.........")
	dburl := os.Getenv("DB_URL_MOCK")
	port := 4000 // any port number will work
	secret := []byte(os.Getenv("SECRET"))
	log.Println(dburl)
	server := &Server{
		port:   port,
		db:     database.NewFrom(dburl),
		secret: secret,
	}
	return server, nil
}

func clearupTesting() {
}
