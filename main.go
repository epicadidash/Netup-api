package main

// import (
// 	"fmt"

// 	parent "dotoday.com/core/database"
// )

// func main() {
// 	parent.Connect()
// 	parent.Ping()
// 	star := parent.SearchUSER(`jr@calhlou.io`, `email`)
// 	hei := parent.UpdateUSER(`starkiller@hei.io`, `email`, star)
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
	r.HandleFunc("/auth/signup", star.ArticleHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
