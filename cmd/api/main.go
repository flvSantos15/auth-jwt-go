package main

import (
	"fmt"
	"net/http"

	"github.com/flvSantos15/auth-jwt-go/internal/server"
)

// Parei em 45:50

func main() {
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
