package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "dotoday.com/core/database"
)

type request struct {
	Dist string `json:"username/email"`
	Pass string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var p request
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if db.TestUser(p.Dist, "email", Hash(p.Pass)) {
		var s response = response{Res: "Succesful"}
		jsonResponse, jsonError := json.Marshal(s)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(jsonResponse)
	} else if db.TestUser(p.Dist, "username", Hash(p.Pass)) {
		var s response = response{Res: "Succesful"}
		jsonResponse, jsonError := json.Marshal(s)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(jsonResponse)
	} else {
		var s response = response{Res: "Unsuccesful"}
		jsonResponse, jsonError := json.Marshal(s)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
	}

}
