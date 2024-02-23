package server

import (
	"clothe-shop-v2/internal/database"
	"clothe-shop-v2/templates"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) AddNewProduct(c *gin.Context) {
	var product database.Product
	var err error
	product.Name = c.PostForm("name")
	product.Description = c.PostForm("description")
	product.Price, err = strconv.Atoi(c.PostForm("price"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "invalid values of price")
		return
	}
	product.ID, err = s.db.AddProduct(product)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to add the product")
		return
	}
	templates.Product(product).Render(context.Background(), c.Writer)
}

		return
	}
	templates.Product(product).Render(context.Background(), c.Writer)
}
