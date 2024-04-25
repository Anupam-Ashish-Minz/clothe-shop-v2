package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type Service interface {
	Health() map[string]string
	GetProducts(page int) ([]Product, error)
	GetProductById(productID int64) (Product, error)
	AddProduct(product Product) (int64, error)
	UpdateProduct(product Product) error
	GetUserByEmail(email string) (User, error)
	GetUserById(userID int64) (User, error)
	AddNewUser(user User) (int64, error)
	ProductsInCart(userID int64) ([]OrderItem, error)
	AddProductInCart(userID int64, productID int64, quantity int) error
	CheckProductInCart(userID int64, productID int64) bool
	UpdateCartProductCount(userID int64, productID int64, increamentQuantity int) error
	GetCartItemById(userID int64, productID int64) (OrderItem, error)
	RemoveCartItem(userID int64, productID int64) error
	GetAllProductsInCart(userID int64) ([]OrderItem, error)
	NewOrder(userID int64, product OrderItem) (int64, error)
	CleanCart(userID int64) error
	GetOrdersFromUser(userID int64) ([]Order, error)
	GetOrderWithProductsFromUser(userID int64) ([]OrderWithProducts, error)
	GetOrderCount(interval OrderCountLength) ([]OrderCount, error)
	GetTotalRevenue(interval OrderCountLength) ([]RevenueAmount, error)
	GetAllOrders() ([]OrderWithProducts, error)
	ChangeOrderStatus(orderID int64, orderStatus OrderStatus) error
	GetAllProducts() ([]Product, error)
}

type service struct {
	db *sqlx.DB
}

var (
	dburl = os.Getenv("DB_URL")
)

func New() Service {
	db, err := sqlx.Connect("postgres", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
