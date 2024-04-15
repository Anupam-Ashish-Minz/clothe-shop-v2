package database

import "time"

type OrderCount struct {
	Date  time.Time
	Count int
}

type OrderCountLength string

const (
	WEEKLY  OrderCountLength = "7 days"
	MONTHLY OrderCountLength = "30 days"
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
