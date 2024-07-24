package server

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"clothe-shop-v2/internal/database"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Product struct {
	ID          int64  `json:"id,omitempty" db:"id"`
	Name        string `json:"name,omitempty" db:"name"`
	Description string `json:"description,omitempty" db:"description"`
	Price       int    `json:"price,omitempty" db:"price"`
	Gender      string `json:"gender,omitempty" db:"gender"`
}

func readCsv() error {
	file, err := os.Open("../../data/csv/products.csv")
	r := csv.NewReader(file)
	products := make([]Product, 0)

	record, err := r.Read()
	if !(record[0] == "name" && record[1] == "description" && record[2] == "price" && record[3] == "gender" && len(record) == 4) {
		return fmt.Errorf("unexpected field names")
	}
	for {
		record, err = r.Read()
		if err != nil {
			break
		}
		var product Product
		product.Name = record[0]
		product.Description = record[1]
		price, err := strconv.ParseFloat(record[2], 32)
		product.Price = int(price)
		if err != nil {
			return err
		}
		product.Gender = record[3]
		products = append(products, product)
	}
	conn := os.Getenv("DB_URL_MOCK")
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return err
	}
	tx := db.MustBegin()
	defer tx.Rollback()
	for _, product := range products {
		_, err = tx.NamedExec(`INSERT INTO "Product" (id, name, description, price, gender) VALUES (:id, :name, :description, :price, :gender)`, product)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	return nil
}

func setupTesting() (*Server, error) {
	godotenv.Load("../../.env")
	dburl := os.Getenv("DB_URL_MOCK")
	db := database.NewFrom(dburl)

	port := 4000 // any port will work
	secret := []byte("mysecret")
	log.Println(dburl)
	server := &Server{
		port:   port,
		db:     db,
		secret: secret,
	}
	return server, nil
}

func clearupTesting() {
}

func connectTestingDatabase() *sqlx.DB {
	godotenv.Load("../../.env")
	dburl := os.Getenv("DB_URL_MOCK")
	db, err := sqlx.Connect("postgres", dburl)
	if err != nil {
		panic(err)
	}
	return db
}
