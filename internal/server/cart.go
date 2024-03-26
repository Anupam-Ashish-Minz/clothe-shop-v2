package server

import (
	"clothe-shop-v2/templates"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) CartPage(c *gin.Context) {
	userID, err := s.Authenticate(c)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	products, err := s.db.ProductsInCart(userID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to fetch products from cart")
		return
	}
	err = templates.Cart(products).Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) AddToCart(c *gin.Context) {
	productID, err := strconv.Atoi(c.PostForm("product_id"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "cannot add product to cart, product id is missing")
		return
	}
	quantity, err := strconv.Atoi(c.PostForm("quantity"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "cannot add product to cart, product id is missing")
		return
	}
	userID, err := s.Authenticate(c)
	if err != nil {
		log.Println(err)
		c.String(http.StatusUnauthorized, "login required to add item to cart")
		return
	}
	if !s.db.CheckProductInCart(userID, int64(productID)) {
		log.Println(err)
		err = s.db.UpdateCartProductCount(userID, int64(productID), quantity)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		log.Println("insert product in cart query")
		err = s.db.AddProductInCart(userID, int64(productID), quantity)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "falied to add data to the cart")
			return
		}
	}
}
