package server

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"clothe-shop-v2/internal/database"
)

type Product struct {
	Name        string
	Description string
	Price       int
	Gender      string
}

func setupTesting() (Server, error) {
	fmt.Println("starting setup.........")
	file, err := os.Open("../../data/csv/products.csv")
	if err != nil {
		return Server{}, err
	}
	r := csv.NewReader(file)
	products := make([]Product, 0)

	for {
		record, err := r.Read()
		if err != nil {
			break
		}
		var product Product
		product.Name = record[0]
		product.Description = record[1]
		product.Price, err = strconv.Atoi(record[2])
		if err != nil {
			return Server{}, err
		}
		product.Gender = record[3]
		products = append(products, product)
	}

	dburl := os.Getenv("DB_URL_TEST")
	db := database.NewFrom(dburl)
	server := Server{
		port:   4001,
		db:     db,
		secret: []byte(os.Getenv("SECRET")),
	}

	return server, err
}

func clearupTesting() {
}
