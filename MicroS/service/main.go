package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Egor"},
	{ID: 2, Name: "Daria"},
}

func getUser(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id not number", http.StatusBadRequest)
	}

	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)

}

func main() {

	http.HandleFunc("/users", getUser)

	log.Println("Сервер запущен")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Не удалось запустить сервер")
	}

}
