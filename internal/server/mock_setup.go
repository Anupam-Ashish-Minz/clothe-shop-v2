package server

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"clothe-shop-v2/internal/database"

	"github.com/joho/godotenv"
)

func readCsv() {
	file, err := os.Open("../../data/csv/products.csv")
	r := csv.NewReader(file)
	products := make([]database.Product, 0)

	record, err := r.Read()
	if !(record[0] == "name" && record[1] == "description" && record[2] == "price" && record[3] == "gender") {
		// handler it later
	}
	for {
		record, err = r.Read()
		if err != nil {
			break
		}
		var product database.Product
		product.Name = record[0]
		product.Description = record[1]
		price, err := strconv.ParseFloat(record[2], 32)
		product.Price = int(price)
		if err != nil {
			// handle it later
		}
		product.Gender = record[3]
		products = append(products, product)
	}
}

func setupTesting() (*Server, error) {
	godotenv.Load("../../.env")
	dburl := os.Getenv("DB_URL_MOCK")
	db := database.NewFrom(dburl)

	port := 4000 // any port will work
	secret := []byte(os.Getenv("SECRET"))
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
