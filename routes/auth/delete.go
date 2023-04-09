package auth

import (
	"fmt"
	"net/http"

	db "dotoday.com/core/database"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.Connect()
	db.RemoveUSER(id)
	// hei := fmt.Sprintf("Sucessfully deleted = %d", star)
	// var s response = response{Res: hei}
	// jsonResponse, jsonError := json.Marshal(s)
	// if jsonError != nil {
	// 	fmt.Println("Unable to encode JSON")
	// }
	fmt.Println(id)
}
