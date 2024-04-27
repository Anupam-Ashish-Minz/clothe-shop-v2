package database

import (
	"time"
)

type OrderStatus string

const (
	STATUS_DELIVERED        OrderStatus = "DELIVERED"
	STATUS_PROCESSING       OrderStatus = "PROCESSING"
	STATUS_CANCLED          OrderStatus = "CANCLED"
	STATUS_OUT_FOR_DELIVERY OrderStatus = "OUT_FOR_DELIVERY"
)

type Order struct {
	ID        int64       `db:"id"`
	Date      time.Time   `db:"date"`
	Status    OrderStatus `db:"status"`
	ProductID int64       `db:"productId"`
	UserID    int64       `db:"userId"`
	Quantity  int         `db:"quantity"`
}

type OrderWithProducts struct {
	ID       int64       `db:"id"`
	Date     time.Time   `db:"date"`
	Status   OrderStatus `db:"status"`
	Product  Product     `db:"product"`
	UserID   int64       `db:"userId"`
	Quantity int         `db:"quantity"`
}

func (s *service) NewOrder(userID int64, product OrderItem) (int64, error) {
	var orderID int64
	err := s.db.Get(&orderID, `insert into "Order" ("productId", "userId", quantity)
		values ($1, $2, $3) returning id`, product.ID, userID, product.Quantity)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func (s *service) GetOrdersFromUser(userID int64) ([]Order, error) {
	var orders []Order
	err := s.db.Select(&orders, `select id, date, state as "status", "productId", "userId",
		quantity from "Order" where "userId" = $1`, userID)
	if err != nil {
		return []Order{}, err
	}
	return orders, nil
}

func (s *service) GetOrderWithProductsFromUser(userID int64) ([]OrderWithProducts, error) {
	var orders []OrderWithProducts
	err := s.db.Select(&orders, `SELECT o.id, o.date, o.state as "status", o."userId",
		o.quantity, p.id as "product.id", p.name as "product.name",
		p.description as "product.description", p.gender as "product.gender",
		p.price as "product.price", p.image as "product.image" FROM "Order" AS o
	INNER JOIN "Product" AS p ON o."productId" = p.id WHERE o."userId" = $1
	ORDER BY o.date DESC`, userID)

	if err != nil {
		return []OrderWithProducts{}, err
	}
	return orders, nil
}
