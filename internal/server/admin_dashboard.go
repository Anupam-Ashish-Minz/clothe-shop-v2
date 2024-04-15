package server

import (
	"clothe-shop-v2/internal/database"
	"clothe-shop-v2/internal/sharedtypes"
	"clothe-shop-v2/templates"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) AdminPage(c *gin.Context) {
	orderCountStats, err := s.db.GetOrderCount(database.ORDER_WEEKLY)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to query database")
		return
	}
	graph := sharedtypes.AdminGraphs{}
	graph.OrderCount.Labels = make([]string, 0)
	graph.OrderCount.Data = make([]int, 0)
	for _, ord := range orderCountStats {
		sDate := strings.Split(fmt.Sprint(ord.Date), " ")
		graph.OrderCount.Labels = append(graph.OrderCount.Labels, sDate[0])
		graph.OrderCount.Data = append(graph.OrderCount.Data, ord.Count)
	}
	err = templates.AdminPage(graph).Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) ChangeOrderCountGraph(c *gin.Context) {
	orderCountType := c.PostForm("order-count-duration")
	log.Println(orderCountType)
	if orderCountType != string(database.ORDER_WEEKLY) && orderCountType != string(database.ORDER_MONTHLY) {
		log.Println("the duration of the graph specfied is not valid", orderCountType)
		c.String(http.StatusBadRequest, "the duration of the graph specfied is not valid")
		return
	}
	orderCountStats, err := s.db.GetOrderCount(database.ORDER_WEEKLY)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to query database")
		return
	}
	orderCount := sharedtypes.Graph{}
	orderCount.Labels = make([]string, 0)
	orderCount.Data = make([]int, 0)
	for _, ord := range orderCountStats {
		sDate := strings.Split(fmt.Sprint(ord.Date), " ")
		orderCount.Labels = append(orderCount.Labels, sDate[0])
		orderCount.Data = append(orderCount.Data, ord.Count)
	}
	templates.OrderCountGraph(orderCount).Render(context.Background(), c.Writer)
}
