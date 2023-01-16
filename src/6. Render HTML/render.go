package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var filepath = path.Join("Render HTMl", "indexfe.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Learning Golang Web",
			"name":  "Nathania Michelle",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("asset"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
