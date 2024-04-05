package database

import (
	"fmt"
	"log"
)

func (s *service) CheckProductInCart(userID int64, productID int64) bool {
	row := s.db.QueryRow(`select count(*) from Cart where userId = ? and productId = ?`,
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
	_, err := s.db.Exec(`insert into Cart (userId, productId, quantity) values (?, ?, ?)`,
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
	row := tx.QueryRow(`select quantity from Cart where userId = ? and productId = ?`, userID, productID)
	var prevQuantity int
	err = row.Scan(&prevQuantity)
	if err != nil {
		return err
	}
	if prevQuantity+incrementQuantity <= 0 {
		_, err = tx.Exec(`delete from Cart where userId = ? and productId = ?`, userID, productID)
		if err != nil {
			return err
		}
		tx.Commit()
	} else {
		_, err = tx.Exec(`update Cart set quantity = ? where userId = ? and productId = ?`,
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
	row := s.db.QueryRow(`select productId, p.name, p.description, p.gender,
		p.price, p.image, quantity from Cart inner join Product as p on productId = p.id
		where userId = ? and productId = ?`, userID, productID)
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
	_, err := s.db.Exec(`delete from Cart where userId = ? and productId = ?`, userID, productID)
	return err
}
