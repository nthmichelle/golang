package main

import (
	"net/http"
	"text/template"
)

type M map[string]interface{}

func main() {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

http.HandleFunc("/about", func(w, http.ResponseWriter, r *http.Request){
	var data = M{"name": "batman"}
	err = tmpl.ExecuteTemplate(w,"index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
})
}
