package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func handleTime(res http.ResponseWriter, req *http.Request) {
	s := time.Now().Format("02.01.2006 15:04:05")
	res.Write([]byte(s))
}

func handleMain(res http.ResponseWriter, req *http.Request) {
	s := fmt.Sprintf("Method: %s\nHost: %s\nPath: %s",
		req.Method, req.Host, req.URL.Path)
	res.Write([]byte(s))
}

func main() {
	r := chi.NewRouter()

	r.Get("/time", handleTime)
	r.Get("/", handleMain)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

	// mux := http.NewServeMux()
	// mux.HandleFunc("/time", handleTime)
	// mux.HandleFunc("/", handleMain)

	// err := http.ListenAndServe(":8080", mux)
	// if err != nil {
	// 	panic(err)
	// }
}
