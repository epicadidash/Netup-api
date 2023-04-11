package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "dotoday.com/core/database"
	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	value := r.URL.Query().Get("value")
	db.Connect()
	vars := mux.Vars(r)
	hei := vars["type"]
	grep := db.UpdateUSER(value, hei, id)
	// fmt.Println(grep)
	hei = fmt.Sprintf(`Succesful = %d`, grep)
	var s response = response{Res: hei}
	jsonResponse, jsonError := json.Marshal(s)
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResponse)
}
