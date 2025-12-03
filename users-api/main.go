package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var users = []string{"Eduard", "Jenny", "Kate"}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 0 || id >= len(users) {
		http.Error(w, "Некорректный идентификатор", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(users[id]))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Имя не определено", http.StatusBadRequest)
		return
	}
	users = append(users, name)
	str := fmt.Sprintf("user %s, id %d", name, len(users)-1)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(str))
}

func main() {
	r := chi.NewRouter()
	r.Get("/user/{id}", getUser)
	r.Get("/users", getUsers)
	r.Post("/user", postUser)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка запуска сервера : %s", err.Error())
		return
	}
}
