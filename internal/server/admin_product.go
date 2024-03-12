package server

import (
	"clothe-shop-v2/internal/database"
	"clothe-shop-v2/templates"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	img, imgHeader, err := c.Request.FormFile("image")
	if err != nil {
		log.Println(err)
		return
	}

	writeMultipartImage := func(img multipart.File, imageName string) error {
		sFilenames := strings.Split(imageName, ".")
		fileExt := sFilenames[len(sFilenames)-1]
		if fileExt != "png" && fileExt != "jpeg" && fileExt != "jpg" {
			return fmt.Errorf("file extension not found")
		}
		filename := "./data/images/" + uuid.New().String() + "." + fileExt

		buf, err := io.ReadAll(img)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, buf, 0644)
		if err != nil {
			return err
		}
		return nil
	}

	err = writeMultipartImage(img, imgHeader.Filename)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to write the image")
	}

	// product.ID, err = s.db.AddProduct(product)
	// if err != nil {
	// 	log.Println(err)
	// 	c.String(http.StatusInternalServerError, "failed to add the product")
	// 	return
	// }
	// templates.Product(product).Render(context.Background(), c.Writer)
}

func (s *Server) UpdateProduct(c *gin.Context) {
	var product database.Product
	var err error
	product.Name = c.PostForm("name")
	product.Description = c.PostForm("description")
	product.Price, err = strconv.Atoi(c.PostForm("price"))
	if err != nil {
		log.Println(err)
		log.Println("invalid value of price setting default price to 0")
		product.Price = 0
	}
	err = s.db.UpdateProduct(product)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "failed to update the product")
		return
	}
	templates.Product(product).Render(context.Background(), c.Writer)
}
