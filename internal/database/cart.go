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

func (s *service) UpdateCartProductCount(userID int64, productID int64, quantity int) error {
	if quantity == 0 {
		return fmt.Errorf("quantity cannot be zero")
	}
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	row := tx.QueryRow(`select quantity from Cart where userId = ? and productId = ?`, userID, productID)
	var prevQuantity int
	err = row.Scan(&prevQuantity)
	if err != nil {
		return err
	}
	tx.Exec(`update Cart set quantity = ? where userId = ? and productId = ?`,
		prevQuantity+quantity, userID, productID)
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
