package database

import (
	"log"
	"time"
)

type OrderStatus string

const (
	COMPLETED OrderStatus = "COMPLETED"
	PENDING   OrderStatus = "PENDING"
	CANCLED   OrderStatus = "CANCLED"
)

type Order struct {
	ID        int64
	Date      time.Time
	Status     OrderStatus
	ProductID int64
	UserID    int64
	Quantity  int
}

func (s *service) NewOrder(userID int64, product OrderItem) (int64, error) {
	row := s.db.QueryRow(`insert into "Order" ("productId", "userId", quantity)
		values ($1, $2, $3) returning id`, product.ID, userID, product.Quantity)
	var orderID int64
	if err := row.Scan(&orderID); err != nil {
		return 0, err
	}
	return orderID, nil
}

func (s *service) GetOrdersFromUser(userID int64) ([]Order, error) {
	rows, err := s.db.Query(`select id, date, state, "productId", "userId",
		quantity from "Order" where "userId" = $1`, userID)
	if err != nil {
		return []Order{}, err
	}

	var orders []Order
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.Date, &order.Status, &order.ProductID,
			&order.UserID, &order.Quantity)
		if err != nil {
			log.Println(err)
		}
		orders = append(orders, order)
	}
	return orders, nil
}
