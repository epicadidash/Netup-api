package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "dotoday.com/core/database"
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
type response struct {
	Res string `json:"result`
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	var p person
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Connect()
	star := fmt.Sprintf("email=%s ", p.Email)
	kei := fmt.Sprintf("username=%s ", p.Username)

	if db.SearchUSER(star) != -1 || db.SearchUSER(kei) != -1 {
		var s response = response{Res: "Bad request"}
		jsonResponse, jsonError := json.Marshal(s)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
	} else {
		var s response = response{Res: "Succesful"}
		jsonResponse, jsonError := json.Marshal(s)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(jsonResponse)
	}
}
