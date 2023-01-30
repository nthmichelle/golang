package main

import "net/http"

const USERNAME = "nathania"
const PASSWORD = "123456"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	usename, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`Something went wrong`))
		return false
	}

	isValid := (usename == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "Get" {
		w.Write([]byte("Only Get Is Allowed"))
		return false

	}
	return true

}
