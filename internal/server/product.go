package server

import (
	"clothe-shop-v2/templates"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) FetchProducts(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Println(err)
		page = 0
	}
	products, err := s.db.GetProducts(page)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to fetch products")
		return
	}
	if len(products) < 1 {
		c.String(http.StatusBadRequest, "there are no products in this page")
		return
	}
	err = templates.Products(products, page+1).Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to load products")
		return
	}
}

func (s *Server) ProductsPage(c *gin.Context) {
	products, err := s.db.GetProducts(0)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to fetch products")
		return
	}
	if len(products) < 1 {
		c.String(http.StatusBadRequest, "no products found")
		return
	}
	err = templates.ProductsPage(products).Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "unable to render template")
	}
}

func (s *Server) ProductPage(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	product, err := s.db.GetProductById(int64(productID))
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to fetch products")
		return
	}
	log.Println(product)
	// err = templates.Product(product).Render(context.Background(), c.Writer)
	// if err != nil {
	// 	log.Println(err)
	// }
}
