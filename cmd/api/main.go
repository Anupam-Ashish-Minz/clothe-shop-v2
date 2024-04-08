package main

import (
	"clothe-shop-v2/internal/server"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--help" {
			return
		}
		return
	}

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
