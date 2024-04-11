package server

import (
	"clothe-shop-v2/templates"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) OrderPage(c *gin.Context) {
	userID, err := s.Authenticate(c)
	if err != nil {
		log.Println(err)
		c.String(http.StatusUnauthorized, "failed to authenticate user")
		return
	}
	// orders, err := s.db.GetOrdersFromUser(userID)
	orders, err := s.db.GetOrderWithProductsFromUser(userID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "order query falied")
		return
	}
	if err := templates.OrderPage(orders).Render(context.Background(), c.Writer); err != nil {
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
	var anyError error = nil
	for _, product := range products {
		orderID, err := s.db.NewOrder(userID, product)
		log.Println(orderID)
		if err != nil {
			log.Println(err)
			anyError = err
		}
	}
	if anyError != nil {
		c.String(http.StatusInternalServerError, "failed to place certain orders, check arrordingly")
		return
	}
	err = s.db.CleanCart(userID)
	if err != nil {
		log.Println(err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/orders")
}
