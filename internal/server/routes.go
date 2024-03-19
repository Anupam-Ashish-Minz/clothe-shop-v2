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
