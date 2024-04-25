package database

import (
	"fmt"
	_ "github.com/lib/pq"
)

type Product struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Gender      string `db:"gender"`
	Price       int    `db:"price"`
	Image       string `db:"image"`
}

type OrderItem struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       int    `db:"price"`
	Gender      string `db:"gender"`
	Image       string `db:"image"`
	Quantity    int    `db:"quantity"`
}

func (s *service) GetProducts(page int) ([]Product, error) {
	pageSize := 10
	var products []Product
	err := s.db.Select(
		&products,
		`select id, name, description, price, gender, image from 
		"Product" limit $1 offset $2`,
		pageSize, page*pageSize,
	)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *service) AddProduct(product Product) (int64, error) {
	if product.Name == "" || product.Price == 0 || product.Gender == "" || product.Image == "" {
		return 0, fmt.Errorf(fmt.Sprint("empty fields in product struct", product))
	}
	row := s.db.QueryRow(`insert into "Product" (name, price, description, gender, image) values ($1, $2, $3, $4, $5) returning id`,
		product.Name, product.Price, product.Description, product.Gender, product.Image)
	var productID int64
	err := row.Scan(&productID)
	if err != nil {
		return 0, err
	}
	return productID, nil
}

func (s *service) GetProductById(productID int64) (Product, error) {
	var product Product
	err := s.db.Get(&product, `select id, name, price, description, gender, image from "Product" where id = $1`, productID)
	if err != nil {
		return Product{}, err
	}
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
	var products []OrderItem
	err := s.db.Select(&products, `select p.id, p.name, p.description, p.price,
		p.gender, p.image, c.quantity from "User" as u join "Cart" as c on u.id =
		c."userId" join "Product" as p on c."productId" = p.id where c."userId" = $1`,
		userID)
	if err != nil {
		return products, err
	}
	return products, nil
}
