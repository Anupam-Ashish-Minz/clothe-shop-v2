package server

import (
	"clothe-shop-v2/templates"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) OrderPage(c *gin.Context) {
	if err := templates.OrderPage().Render(context.Background(), c.Writer); err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to render template")
		return
	}
}

func (s *Server) PlaceOrder(c *gin.Context) {
	userID, err := s.Authenticate(c)
	if err != nil {
		log.Println(err)
		c.String(http.StatusUnauthorized, "login required")
		return
	}
	products, err := s.db.GetAllProductsInCart(userID)
	if err != nil {
		log.Println(err)
	}
	log.Println(products)
}
