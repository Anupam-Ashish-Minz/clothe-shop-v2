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
