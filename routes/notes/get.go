package notes

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "dotoday.com/core/database"
)

type response2 struct {
	Hello   []db.Note `json:"response"`
	Message string    `json:"message"`
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	db.Connect()

	grep := db.GetNote(id)
	// fmt.Println(grep)
	var s response2 = response2{Hello: grep, Message: `Successful`}
	jsonResponse, jsonError := json.Marshal(s)
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResponse)
}
