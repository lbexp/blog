package main

import (
	"fmt"
	core_http "lbexp-blog/core/http"
	"net/http"
	"os"
)

func main() {
	server := core_http.NewServer()

	server.Router.GET("/blogs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Blogs get all")
	})
	server.Router.POST("/blogs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Blogs post")
	})
	server.Router.GET("/blogs/:id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Blogs get by id")
	})
	server.Router.PUT("/blogs/:id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Blogs update by id")
	})
	server.Router.DELETE("/blogs/:id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Blogs delete by id")
	})

	server.Serve()

	os.Exit(1)
}
