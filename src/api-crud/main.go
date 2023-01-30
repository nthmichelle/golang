package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGET(w, r) {
		return
	}
	if id := r.URL.Query().Get("id"); id != "" {

		OutputJSON(w, GetStudents())
	}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	ress, err := Json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

const USERNAME = "nathania"
const PASSWORD = "123456"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "Get" {
		w.Write([]byte("Only GET Is Allowed"))
		return false
	}

	return true
}
}
