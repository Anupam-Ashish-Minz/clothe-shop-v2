package server

import (
	"clothe-shop-v2/templates"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.StaticFile("/static/output.css", "./templates/output.css")
	r.StaticFile("/static/htmx.min.js", "./static/htmx.min.js")
	r.StaticFile("/static/alpine.min.js", "./static/alpine.min.js")
	r.Static("/static/images/", "./data/images")

	r.GET("/", s.HomePage)
	r.GET("/health", s.healthHandler)

	r.GET("/login", func(c *gin.Context) {
		templates.Login().Render(context.Background(), c.Writer)
	})
	r.GET("/signup", func(c *gin.Context) {
		templates.Signup().Render(context.Background(), c.Writer)
	})
	r.POST("/login", s.UserLogin)
	r.POST("/signup", s.UserSignup)

	r.GET("/products", s.ProductsPage)
	r.GET("/api/products/", s.FetchProducts)
	r.GET("/product/:id", s.ProductPage)

	r.GET("/cart", s.CartPage)
	r.POST("/api/cart/add/", s.AddToCart)

	r.GET("/admin", s.AdminPage)

	r.POST("/admin/product/add", s.AddNewProduct)
	r.POST("/admin/product/update", s.UpdateProduct)

	return r
}

func (s *Server) HomePage(c *gin.Context) {
	err := templates.Index().Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) AdminPage(c *gin.Context) {
	err := templates.AdminPage().Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
	}
}
