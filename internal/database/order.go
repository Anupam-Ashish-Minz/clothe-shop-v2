package database

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
