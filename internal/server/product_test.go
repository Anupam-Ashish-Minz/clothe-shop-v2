package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func testProducts(router http.Handler, t *testing.T) {
// 	r := httptest.NewRequest("GET", "http://127.0.0.1:4000/products", nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, r)
// 	router.ServeHTTP(w, r)
// 	if w.Code != http.StatusOK {
// 		data, err := io.ReadAll(w.Body)
// 		if err != nil {
// 			t.Log("failed to get the renspone body")
// 			t.Log(err)
// 		}
// 		log.Println("product page:", string(data))
// 		t.Fatal("failed to fetch product page, with status code", w.Code)
// 	}
// }
//
// func TestProductRoutes(t *testing.T) {
// 	s, err := setupTesting()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	router := s.RegisterRoutes()
//
// 	testProducts(router, t)
// }

func TestNullCheckProduct(t *testing.T) {
	assert := assert.New(t)
	p1, err := CheckProduct("mock product #1", "male", "this is a mock product", "5000")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(p1.Name, "mock product #1")
	assert.Equal(p1.Gender, "male")
	assert.Equal(p1.Description, "this is a mock product")
	assert.Equal(p1.Price, 5000)
}
