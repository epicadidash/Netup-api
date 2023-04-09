package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "dotoday.com/core/database"
)

type person struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Email      string `json:"email"`
	Last_name  string `json:"last_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
}
type response struct {
	Res string `json:"result`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var p person
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Connect()

	if db.SearchUSER(p.Username, `username`) == -1 {
		star := db.SearchUSER(p.Username, `username`)
		hei := fmt.Sprintf("Bad = %d", star)
		var s response = response{Res: hei}
		jsonResponse, jsonError := json.Marshal(s)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
	} else if db.SearchUSER(p.Email, `email`) == -1 {
		star := db.SearchUSER(`jr@calhlou.io`, `email`)
		hei := fmt.Sprintf("Bad = %d", star)
		var s response = response{Res: hei}
		jsonResponse, jsonError := json.Marshal(s)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
	} else {
		db.InsertUSER(p.Email, p.First_name, p.Last_name, p.Username, Hash(p.Password))
		var s response = response{Res: "Succesful"}
		jsonResponse, jsonError := json.Marshal(s)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(jsonResponse)
	}
}
