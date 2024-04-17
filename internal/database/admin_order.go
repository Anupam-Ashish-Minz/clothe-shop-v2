package database

import "time"

type OrderCount struct {
	Date  time.Time
	Count int
}

type RevenueAmount struct {
	Date   time.Time
	Amount int
}

type OrderCountLength string

const (
	ORDER_WEEKLY  OrderCountLength = "7 days"
	ORDER_MONTHLY OrderCountLength = "30 days"
)

func (s *service) GetOrderCount(interval OrderCountLength) ([]OrderCount, error) {
	rows, err := s.db.Query(`select date::date, count(*) from "Order" where
		date::date > now() - $1::interval group by date::date order by date::date`, interval)
	if err != nil {
		return []OrderCount{}, err
	}
	orderCounts := make([]OrderCount, 0)
	for rows.Next() {
		var orderCount OrderCount
		err = rows.Scan(&orderCount.Date, &orderCount.Count)
		if err != nil {
			return []OrderCount{}, err
		}
		orderCounts = append(orderCounts, orderCount)
	}
	return orderCounts, nil
}

func (s *service) GetTotalRevenue(interval OrderCountLength) ([]RevenueAmount, error) {
	rows, err := s.db.Query(`select date::date, sum(quantity * p.price) from
		"Order" as o inner join "Product" as p on o."productId" = p.id where
		date::date > now() - $1::interval group by date::date order by
		date::date`, interval)
	if err != nil {
		return []RevenueAmount{}, err
	}
	revenueAmounts := make([]RevenueAmount, 0)
	for rows.Next() {
		var revenueAmount RevenueAmount
		err = rows.Scan(&revenueAmount.Date, &revenueAmount.Amount)
		if err != nil {
			return []RevenueAmount{}, err
		}
		revenueAmounts = append(revenueAmounts, revenueAmount)
	}
	return revenueAmounts, nil
}

func (s *service) GetAllOrders() ([]OrderWithProducts, error) {
	rows, err := s.db.Query(`select o.id, date, state, quantity, p.id, p.name,
		p.price, p.description, p.gender, p.image from "Order" as o inner join
		"Product" as p on o."productId" = p.id`)
	if err != nil {
		return []OrderWithProducts{}, err
	}

	var orders []OrderWithProducts
	for rows.Next() {
		var order OrderWithProducts
		err = rows.Scan(&order.ID, &order.Date, &order.Status, &order.Quantity,
			&order.Product.ID, &order.Product.Name, &order.Product.Price,
			&order.Product.Description, &order.Product.Gender,
			&order.Product.Image)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (s *service) ChangeOrderStatus(orderID int64, orderStatus OrderStatus) error {
	_, err := s.db.Exec(`update "Order" set state = $1 where id = $2`, orderStatus, orderID)
	if err != nil {
		return err
	}
	return nil
}
