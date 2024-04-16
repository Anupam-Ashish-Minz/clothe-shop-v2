package server

import (
	"clothe-shop-v2/templates"
	"context"
	"log"
	"net/http"

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
