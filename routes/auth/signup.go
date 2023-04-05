package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	Id         int    `json:"id"`
	First_name string `json:"email"`
	Email      string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	var p person
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(p)
	jsonResponse, jsonError := json.Marshal(p)
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
