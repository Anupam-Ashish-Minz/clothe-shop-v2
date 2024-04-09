package database

import (
	"fmt"
	"log"
)

func (s *service) CheckProductInCart(userID int64, productID int64) bool {
	row := s.db.QueryRow(`select count(*) from "Cart" where "userId" = $1 and "productId" = $2`,
		userID, productID)
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println(err)
		return false
	}
	if count == 0 {
		return true
	} else {
		return false
	}
}

func (s *service) AddProductInCart(userId int64, productID int64, quantity int) error {
	_, err := s.db.Exec(`insert into "Cart" ("userId", "productId", quantity) values ($1, $2, $3)`,
		userId, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateCartProductCount(userID int64, productID int64, incrementQuantity int) error {
	if incrementQuantity == 0 {
		return fmt.Errorf("quantity cannot be zero")
	}
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRow(`select quantity from "Cart" where "userId" = $1 and "productId" = $2`, userID, productID)
	var prevQuantity int
	err = row.Scan(&prevQuantity)
	if err != nil {
		return err
	}
	if prevQuantity+incrementQuantity > 0 {
		_, err = tx.Exec(`update "Cart" set quantity = $1 where "userId" = $2 and "productId" = $3`,
			prevQuantity+incrementQuantity, userID, productID)
		if err != nil {
			return err
		}
		err = tx.Commit()
	}
	return nil
}

func (s *service) GetCartItemById(userID int64, productID int64) (OrderItem, error) {
	var product OrderItem
	row := s.db.QueryRow(`select "productId", p.name, p.description, p.gender,
		p.price, p.image, quantity from "Cart" inner join "Product" as p on "productId" = p.id
		where "userId" = $1 and "productId" = $2`, userID, productID)
	err := row.Scan(&product.ID, &product.Name, &product.Description,
		&product.Gender, &product.Price, &product.Image, &product.Quantity)
	if err != nil {
		return OrderItem{}, err
	}
	return product, nil
}

func (s *service) RemoveCartItem(userID int64, productID int64) error {
	if userID == 0 || productID == 0 {
		return fmt.Errorf("invalid user id or product id")
	}
	_, err := s.db.Exec(`delete from "Cart" where userId = $1 and productId = $2`, userID, productID)
	return err
}

func (s *service) GetAllProductsInCart(userID int64) ([]OrderItem, error) {
	rows, err := s.db.Query(`select "productId", p.name, p.description, p.gender,
		p.price, p.image, quantity from "Cart" inner join "Product" as p on "productId" = p.id
		where "Cart"."userId" = $1`, userID)

	products := make([]OrderItem, 0)
	for rows.Next() {
		var product OrderItem
		err = rows.Scan(&product.ID, &product.Name, &product.Description,
			&product.Gender, &product.Price, &product.Image, &product.Quantity)
		products = append(products, product)
		if err != nil {
			return []OrderItem{}, err
		}
	}
	if err != nil {
		return []OrderItem{}, err
	}
	return products, nil
}

func (s *service) CleanCart(userID int64) error {
	_, err := s.db.Exec(`delete from "Cart" where "userId" = $1`, userID)
	if err != nil {
		return err
	}
	return nil
}
