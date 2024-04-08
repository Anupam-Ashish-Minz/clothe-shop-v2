package database

import (
	"fmt"
	_ "github.com/lib/pq"
)

type Product struct {
	ID          int64
	Name        string
	Description string
	Gender      string
	Price       int
	Image       string
}

type OrderItem struct {
	ID          int64
	Name        string
	Description string
	Price       int
	Gender      string
	Image       string
	Quantity    int
}

func (s *service) GetProducts(page int) ([]Product, error) {
	pageSize := 10
	products := make([]Product, 0)
	var p Product
	rows, err := s.db.Query(
		`select id, name, description, price, gender, image from 
		"Product" limit $1 offset $2`,
		pageSize, page*pageSize,
	)
	if err != nil {
		return products, err
	}
	for rows.Next() {
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Gender, &p.Image)
		products = append(products, p)
	}
	return products, nil
}

func (s *service) AddProduct(product Product) (int64, error) {
	if product.Name == "" || product.Price == 0 || product.Gender == "" || product.Image == "" {
		return 0, fmt.Errorf(fmt.Sprint("empty fields in product struct", product))
	}
	res, err := s.db.Exec(`insert into "Product" (name, price, description, gender, image) values ($1, $2, $3, $4, $5)`,
		product.Name, product.Price, product.Description, product.Gender, product.Image)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (s *service) GetProductById(productID int64) (Product, error) {
	var product Product
	row := s.db.QueryRow(`select id, name, price, description, gender, image from "Product" where id = $1`, productID)
	row.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Gender, &product.Image)
	if product.ID == 0 {
		return product, fmt.Errorf("no such product exists")
	}
	return product, nil
}

func (s *service) UpdateProduct(product Product) error {
	if product.ID == 0 {
		return fmt.Errorf("failed to find product by id")
	}
	orignalProduct, err := s.GetProductById(product.ID)
	if err != nil {
		return err
	}
	if product.Name == "" {
		product.Name = orignalProduct.Name
	}
	if product.Description == "" {
		product.Description = orignalProduct.Description
	}
	if product.Price == 0 {
		product.Price = orignalProduct.Price
	}
	if product.Image == "" {
		product.Image = orignalProduct.Image
	}
	if product.Gender == "" {
		product.Gender = orignalProduct.Gender
	}
	s.db.Exec(`update "Product" set name = $1, description = $2, price = $3, gender = $4, image = $5 where id = $6`,
		product.Name, product.Description, product.Price, product.ID)
	return fmt.Errorf("not implemented")
}

func (s *service) ProductsInCart(userID int64) ([]OrderItem, error) {
	products := make([]OrderItem, 0)
	var product OrderItem
	rows, err := s.db.Query(`select p.id, p.name, p.description, p.price,
		p.gender, p.image, c.quantity from "User" as u join "Cart" as c on u.id =
		c."userId" join "Product" as p on c."productId" = p.id where c."userId" = $1`,
		userID)
	if err != nil {
		return products, err
	}
	for rows.Next() {
		rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Gender, &product.Image, &product.Quantity)
		products = append(products, product)
	}
	return products, nil
}
