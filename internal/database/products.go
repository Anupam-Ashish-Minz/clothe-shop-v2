package database

import "fmt"

type Product struct {
	ID          int64
	Name        string
	Description string
	Gender      string
	Price       int
	Image       string
}

func (s *service) GetProducts() ([]Product, error) {
	products := make([]Product, 0)
	var p Product
	rows, err := s.db.Query(`select id, name, description, price, gender, image from Product`)
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
	res, err := s.db.Exec(`insert into Product (name, price, description, gender, image) values (?, ?, ?, ?, ?)`,
		product.Name, product.Price, product.Description, product.Gender, product.Image)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (s *service) GetProductById(productID int64) (Product, error) {
	var product Product
	row := s.db.QueryRow(`select id, name, price, description, gender, image from Product where id = ?`, productID)
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
	s.db.Exec(`update Product set name = ?, description = ?, price = ? where id = ?`,
		product.Name, product.Description, product.Price, product.ID)
	return fmt.Errorf("not implemented")
}

func (s *service) ProductsInCart(userID int64) ([]Product, error) {
	products := make([]Product, 0)
	var product Product
	rows, err := s.db.Query(`select id, name, description, price, gender from products`)
	if err != nil {
		return products, err
	}
	for rows.Next() {
		rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Gender)
		products = append(products, product)
	}
	return products, nil
}
