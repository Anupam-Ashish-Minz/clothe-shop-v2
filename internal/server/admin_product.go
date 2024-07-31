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

func CheckProduct(name, gender, description, price string) (database.Product, error) {
	if name == "" || gender == "" {
		return database.Product{}, fmt.Errorf("name or gender missing from product")
	}
	intPrice, err := strconv.Atoi(price)
	if err != nil {
		fprice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			return database.Product{}, err
		}
		intPrice = int(fprice)
	}
	return database.Product{
		Name:        name,
		Gender:      gender,
		Description: description,
		Price:       intPrice,
	}, nil
}

func (s *Server) AddNewProduct(c *gin.Context) {
	name := c.PostForm("name")
	gender := c.PostForm("gender")
	description := c.PostForm("description")
	price := c.PostForm("price")

	product, err := CheckProduct(name, gender, description, price)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "invalid values of price")
		return
	}
	img, imgHeader, err := c.Request.FormFile("image")
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "image not found")
		return
	}

	writeMultipartImage := func(img multipart.File, imgHeader *multipart.FileHeader) (string, error) {
		sFilenames := strings.Split(imgHeader.Filename, ".")
		fileExt := sFilenames[len(sFilenames)-1]
		if fileExt != "png" && fileExt != "jpeg" && fileExt != "jpg" {
			return "", fmt.Errorf("file extension not found")
		}
		outFilename := uuid.New().String() + "." + fileExt

		buf, err := io.ReadAll(img)
		if err != nil {
			return outFilename, err
		}
		err = os.WriteFile("./data/images/"+outFilename, buf, 0644)
		if err != nil {
			return outFilename, err
		}
		return outFilename, nil
	}

	imageName, err := writeMultipartImage(img, imgHeader)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to write the image")
		return
	}
	product.Image = imageName

	product.ID, err = s.db.AddProduct(product)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to add the product")
		return
	}
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
	// templates.Product(product).Render(context.Background(), c.Writer)
}

func (s *Server) AdminProductPage(c *gin.Context) {
	products, err := s.db.GetAllProducts()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to fetch products")
		return
	}
	err = templates.AdminProductPage(products).Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to parse the template")
		return
	}
}

func (s *Server) AdminNewProductPage(c *gin.Context) {
	err := templates.AdminAddProductPage().Render(context.Background(), c.Writer)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to parse the template")
	}
}
