package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	s.db.AddProductToCart(int64(productID), quantity)
}
