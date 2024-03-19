package database

func (s *service) AddProductInCart(userId int64, productID int64, quantity int) error {
	_, err := s.db.Exec(`insert into Cart (userId, productId, quantity) values (?, ?, ?)`,
		userId, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}
