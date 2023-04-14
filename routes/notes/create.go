package notes

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "dotoday.com/core/database"
)

type response struct {
	Res string `json:"result`
}
type Note struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Userid      int    `json:"userid"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var p Note
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Connect()
	db.InsertNote(p.Title, p.Description, p.Userid)
	var s response = response{Res: "Succesful"}
	jsonResponse, jsonError := json.Marshal(s)
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResponse)
}
