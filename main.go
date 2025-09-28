package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()
	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Base path")
	})
	router.GET("/for/:id/demonstration/:otherid", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Path with param id")
	})

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}

	fmt.Println("Server started on ", l.Addr().String())

	if err := http.Serve(l, router); err != nil {
		fmt.Printf("Server closed: %s\n", err)
	}

	os.Exit(1)
}
