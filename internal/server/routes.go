package server

import (
	"clothe-shop-v2/internal/database"
	"clothe-shop-v2/templates"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.StaticFile("/static/output.css", "./templates/output.css")
	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	var products []database.Product
	templates.Index(products).Render(context.Background(), c.Writer)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
