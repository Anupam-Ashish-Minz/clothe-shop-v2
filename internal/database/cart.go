package database

func (s *service) AddProductInCart(productID int64, quantity int) error {
	// s.db.Exec(`insert into Cart (product_id, quantity) values (?, ?)`, productID, quantity)
	return nil
}
