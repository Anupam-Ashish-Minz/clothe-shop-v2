package server

import (
	"clothe-shop-v2/internal/database"
	"clothe-shop-v2/templates"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) AdminOrderPage(c *gin.Context) {
	orders, err := s.db.GetAllOrders()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to fetch orders")
		return
	}
	templates.AdminOrders(orders).Render(context.Background(), c.Writer)
}

func (s *Server) AdminChangeOrderStatus(c *gin.Context) {
	i32orderID, err := strconv.Atoi(c.Param("order_id"))
	orderID := int64(i32orderID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "invalid order id")
		return
	}
	orderStatus := c.Param("order_status")
	if orderStatus != string(database.STATUS_DELIVERED) &&
		orderStatus != string(database.STATUS_OUT_FOR_DELIVERY) &&
		orderStatus != string(database.STATUS_CANCLED) {
		log.Println("invalid order status")
		c.String(http.StatusBadRequest, "invaild order status")
		return
	}
	err = s.db.ChangeOrderStatus(int64(orderID), database.OrderStatus(orderStatus))
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to process the database query")
		return
	}

	order, err := s.db.GetOrderByID(orderID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "unable to find the order")
		return
	}
	orderWithProducts := database.OrderWithProducts{
		ID:   order.ID,
		Date: order.Date,
	}
	err = templates.OrderStatusCell().Render(context.Background(), c.Writer)

	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "unable to render order template")
		return
	}
}
