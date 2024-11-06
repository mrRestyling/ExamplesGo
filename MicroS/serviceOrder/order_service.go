package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Order struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	Total  int `json:"total"`
}

type OrderResponse struct {
	ID    int `json:"id"`
	Total int `json:"total"`
}

type User struct {
	Name string `json:"name"`
}

var orders = []Order{
	{ID: 1, UserID: 1, Total: 100},
	{ID: 2, UserID: 2, Total: 200},
	{ID: 3, UserID: 1, Total: 200},
}

func getUserOrders(w http.ResponseWriter, r *http.Request) {
	userIDstr := r.URL.Query().Get("user_id")

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		http.Error(w, "Неверный формат", http.StatusBadRequest)
	}

	resp, err := http.Get("http://localhost:8080/users?id=1")
	if err != nil {
		http.Error(w, "Неверный формат", http.StatusBadRequest)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userOrders []Order
	for _, order := range orders {
		if order.UserID == userID {
			userOrders = append(userOrders, order)
		}
	}

	var orderResponses []OrderResponse
	for _, order := range userOrders {
		orderResponses = append(orderResponses, OrderResponse{order.ID, order.Total})
	}

	json.NewEncoder(w).Encode("заказы пользователя: " + user.Name)
	json.NewEncoder(w).Encode(orderResponses)

}

func main() {

	http.HandleFunc("/users/orders", getUserOrders)

	http.ListenAndServe(":8081", nil)
}
