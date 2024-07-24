package server

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testProducts(router http.Handler, t *testing.T) {
	r := httptest.NewRequest("GET", "http://127.0.0.1:4000/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	router.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		data, err := io.ReadAll(w.Body)
		if err != nil {
			t.Log("failed to get the renspone body")
			t.Log(err)
		}
		log.Println("product page:", string(data))
		t.Fatal("failed to fetch product page, with status code", w.Code)
	}
}

func TestProductRoutes(t *testing.T) {
	s, err := setupTesting()
	if err != nil {
		t.Fatal(err)
	}
	router := s.RegisterRoutes()

	testProducts(router, t)
}
