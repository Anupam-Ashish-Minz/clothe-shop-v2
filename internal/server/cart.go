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
	err = templates.Cart(products, true).Render(context.Background(), c.Writer)
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

func (s *Server) CartIncreaseProductQuantity(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.String(http.StatusBadRequest, "product id is incorrect")
		return
	}
	userID, err := s.Authenticate(c)
	if err != nil {
		log.Println(err)
		c.String(http.StatusUnauthorized, "login requried to perform action")
		return
	}
	err = s.db.UpdateCartProductCount(userID, int64(productID), 1)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to increamet product count, check if the product is added to the cart")
		return
	}
	products, err := s.db.ProductsInCart(userID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to query data from database")
		return
	}
	err = templates.CartContent(products).Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to render template")
		return
	}
}

func (s *Server) CartDecreaseProductQuantity(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.String(http.StatusBadRequest, "product id is incorrect")
		return
	}
	userID, err := s.Authenticate(c)
	if err != nil {
		log.Println(err)
		c.String(http.StatusUnauthorized, "login requried to perform action")
		return
	}
	err = s.db.UpdateCartProductCount(userID, int64(productID), -1)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to decrement product count, check if the product is count is zero")
		return
	}
	products, err := s.db.ProductsInCart(userID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to query data from database")
		return
	}
	err = templates.CartContent(products).Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to render template")
		return
	}
}

func (s *Server) RemoveItemCart(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.String(http.StatusBadRequest, "product id is incorrect")
		return
	}
	userID, err := s.Authenticate(c)
	if err != nil {
		log.Println(err)
		c.String(http.StatusUnauthorized, "login requried to perform action")
		return
	}
	err = s.db.RemoveCartItem(userID, int64(productID))
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to remove the provided product")
		return
	}
	products, err := s.db.ProductsInCart(userID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to query data from database")
		return
	}
	err = templates.CartContent(products).Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "unable to render template")
		return
	}
}
