package server

import (
	"log"
	"testing"
)

func TestProduct(t *testing.T) {
	_, err := setupTesting()
	if err != nil {
		log.Println(err)
	}
	t.Fatal("fail this test")
}
