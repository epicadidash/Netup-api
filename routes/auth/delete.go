package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "dotoday.com/core/database"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	db.Connect()
	star := db.RemoveUSER(id)

	if star == -1 {
		var s response = response{Res: "Not deleted "}
		jsonResponse, jsonError := json.Marshal(s)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
	} else {
		hei := fmt.Sprintf("Sucessfully deleted = %d", star)
		var s response = response{Res: hei}
		jsonResponse, jsonError := json.Marshal(s)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(jsonResponse)
	}
}
