package database

func (s *service) GetAllProducts() ([]Product, error) {
	rows, err := s.db.Query(`select id, name, description, price, gender, image from "Product"`)
	products := make([]Product, 0)
	if err != nil {
		return []Product{}, err
	}
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Gender, &p.Image)
		if err != nil {
			return products, err
		}
		products = append(products, p)
	}
	return products, nil
}
