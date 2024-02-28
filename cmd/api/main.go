package main

import (
	"clothe-shop-v2/internal/database"
	"clothe-shop-v2/internal/server"
	"os"
	"strconv"
	"strings"
)

func parseCSVProducts(file string) ([]database.Product, error) {
	var products []database.Product
	var err error
	err = nil

	lines := strings.Split(file, "\n")
	headerLine := lines[0]
	lines = lines[1:]
	nameIndex := -1
	genderIndex := -1
	descriptionIndex := -1
	priceIndex := -1

	headers := strings.Split(headerLine, ",")
	for i, header := range headers {
		header = strings.Trim(header, " ")
		if header == "Product" {
			nameIndex = i
		} else if header == "Description" {
			descriptionIndex = i
		} else if header == "Price" {
			priceIndex = i
		} else if header == "Gender" {
			genderIndex = i
		}
	}

	for _, line := range lines {
		// log.Println(scanner.Text())
		line := strings.Split(line, ",")
		var product database.Product

		for i := range line {
			x := strings.Trim(line[i], " ")
			if i == nameIndex {
				product.Name = x
			} else if i == descriptionIndex {
				product.Description = x
			} else if i == priceIndex {
				product.Price, err = strconv.Atoi(x)
				if err != nil {
					return products, err
				}
			} else if i == genderIndex {
				// product.Gender = x
			}
			products = append(products, product)
		}
	}
	return products, nil
}

func importDataCSV(filename string) error {
	file, err := os.ReadFile("./tmp/data/" + filename)
	if err != nil {
		return err
	}
	// scanner := bufio.NewScanner(file)
	// scanner.Scan()
	parseCSVProducts(string(file))
	return nil
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--help" {
			return
		}
		if os.Args[1] == "--import" {
			importDataCSV("products-20.csv")
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
