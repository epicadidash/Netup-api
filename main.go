package main

// import (
// 	"fmt"

// 	parent "dotoday.com/core/database"
// )

// func main() {
// 	parent.Connect()
// 	parent.Ping()
// 	star := parent.SearchUSER(`swati@pani`, `email`)
// 	fmt.Println(star)
// }

// import (
// 	"fmt"

// 	parent "dotoday.com/core/routes/auth"
// )

// func main() {
// 	hei := parent.Hash(`starkiller@hei.io`)
// 	fmt.Println(hei)
// }

import (
	"log"
	"net/http"
	"time"

	star "dotoday.com/core/routes/auth"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/auth/delete", star.Delete).Methods("DELETE")
	r.HandleFunc("/auth/signup", star.Signup).Methods("POST")
	r.HandleFunc("/auth/login", star.Login).Methods("POST")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
