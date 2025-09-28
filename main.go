package main

import (
	"fmt"
	utils_http "lbexp-blog/utils/http"
	"net/http"
	"os"
)

func main() {
	server := utils_http.NewServer()

	server.Router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Base path")
	})
	server.Router.GET("/for/:id/demonstration/:otherid", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Path with param id")
	})

	server.Serve()
	os.Exit(1)
}
