package database

import "time"

type Order struct {
	ID        int64
	Date      time.Time
	productID int64
	userID    int64
	quantity  int
}

func (s *service) NewOrder(userID int64, product OrderItem) (int64, error) {
	res, err := s.db.Exec(`insert into "Order" ("productId", "userId", quantity)
		values ($1, $2, $3)`, product.ID, userID, product.Quantity)
	if err != nil {
		return 0, err
	}
	orderID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func (s *service) GetAllOrdersFromUser(userID int64) (Order, error) {
	return Order{}, nil
}
