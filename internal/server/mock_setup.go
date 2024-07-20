package server

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

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

	file, err := os.Open("../../data/csv/products.csv")
	if err != nil {
		return &Server{}, err
	}
	r := csv.NewReader(file)
	products := make([]Product, 0)

	record, err := r.Read()
	if !(record[0] == "name" && record[1] == "description" && record[2] == "price" && record[3] == "gender") {
		return &Server{}, fmt.Errorf("invalid fields in csv data: %s", record)
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
			return &Server{}, err
		}
		product.Gender = record[3]
		products = append(products, product)
	}

	log.Println(products)

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
