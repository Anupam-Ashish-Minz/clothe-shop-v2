package server

import (
	"clothe-shop-v2/internal/database"
	"clothe-shop-v2/templates"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) AdminPage(c *gin.Context) {
	orderCountStats, err := s.db.GetOrderCount(database.WEEKLY)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to query database")
		return
	}
	log.Println(orderCountStats)
	err = templates.AdminPage().Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
	}
}
