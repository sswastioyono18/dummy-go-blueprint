package main

import (
	"dummy-go-blueprint/internal/server"
	"fmt"
)

func main() {

	server := server.NewServer()
	fmt.Println("test")
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
