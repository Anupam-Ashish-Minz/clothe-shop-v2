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

	return r
}

func (s *Server) HomePage(c *gin.Context) {
	products, err := s.db.GetProducts()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to fetch products")
		return
	}
	templates.Index(products).Render(context.Background(), c.Writer)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
