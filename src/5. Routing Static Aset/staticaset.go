package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("Server Started at Localhost:9000")
	http.ListenAndServe(":9000", nil)
}
