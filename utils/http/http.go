package utils_http

import (
	"fmt"
	"net"
	"net/http"
)

type server struct {
	Router *router
}

func NewServer() *server {
	router := newRouter()

	return &server{Router: router}
}

func (s *server) Serve() {
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}

	fmt.Println("Server started on ", l.Addr().String())

	if err := http.Serve(l, s.Router); err != nil {
		fmt.Printf("Server closed: %s\n", err)
	}
}
