package server

import (
	"clothe-shop-v2/templates"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.StaticFile("/static/output.css", "./templates/output.css")
	r.StaticFile("/static/htmx.min.js", "./static/htmx.min.js")

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
	r.GET("/product/:id", s.ProductPage)

	r.GET("/cart", s.CartPage)

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

func (s *Server) ProductsPage(c *gin.Context) {
	products, err := s.db.GetProducts()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to fetch products")
		return
	}
	err = templates.Products(products).Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
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
	err = templates.Product(product).Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
	}
}

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

func (s *Server) AdminPage(c *gin.Context) {
	err := templates.AdminPage().Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
	}
}
