package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

var AccessToken *MyTokens
var Expiration *time.Time

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/Static/Landing_Page/*", http.StripPrefix("/Static/Landing_Page/", http.FileServer(http.Dir("./Static/Landing_Page/"))))
	r.Handle("/Static/SpotifyPage/*", http.StripPrefix("/Static/SpotifyPage/", http.FileServer(http.Dir("./Static/SpotifyPage/"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./Static/Landing_Page/test_me.html")
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.FormValue("input"))
		AddSong(r.FormValue("input"))
	})

	http.ListenAndServe(":8080", r)
}
