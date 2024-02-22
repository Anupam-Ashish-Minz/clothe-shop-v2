package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type Service interface {
	Health() map[string]string
	GetProducts() ([]Product, error)
	AddProduct(product Product) (int64, error)
}

type service struct {
	db *sql.DB
}

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       int
}

var (
	dburl = os.Getenv("DB_URL")
)

func New() Service {
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) GetProducts() ([]Product, error) {
	products := make([]Product, 0)
	var p Product
	rows, err := s.db.Query(`select id, name, description, price from Product`)
	if err != nil {
		return products, err
	}
	for rows.Next() {
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		products = append(products, p)
	}
	return products, nil
}

func (s *service) AddProduct(product Product) (int64, error) {
	if product.Name == "" || product.Price == 0 {
		return 0, fmt.Errorf("empty name or price for product")
	}
	res, err := s.db.Exec(`insert into products (name, price, description) value (?, ?, ?)`,
		product.Name, product.Price, product.Description)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

