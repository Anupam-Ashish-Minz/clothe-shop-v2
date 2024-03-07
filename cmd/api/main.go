package main

import (
	"clothe-shop-v2/internal/server"
	"database/sql"
	"encoding/csv"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func importDataCSV(filename string) error {
	file, err := os.Open("./tmp/data/" + filename)
	if err != nil {
		return err
	}
	reader := csv.NewReader(file)
	// Product, Gender, Description, Price
	data, err := reader.ReadAll()
	if err != nil {
		return err
	}

	DB_TYPE := "sqlite3"
	DB_NAME := os.Getenv("DB_URL")

	db, err := sql.Open(DB_TYPE, DB_NAME)
	defer db.Close()
	if err != nil {
		return err
	}
	for i, rec := range data {
		if i == 0 {
			continue
		}
		_, err := db.Exec(`insert into Product (name, gender, description, price) values (?, ?, ?, ?)`,
			rec[0], rec[1], rec[2], rec[3])
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--help" {
			return
		}
		if os.Args[1] == "--import" {
			// err := importDataCSV("products-20.csv")
			// if err != nil {
			// 	log.Println(err)
			// }
			return
		}
		return
	}

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
