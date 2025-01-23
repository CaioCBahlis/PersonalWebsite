package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

var AccessToken *MyTokens
var Expiration *time.Time

type InputData struct {
	Input string `json:"input"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/Static/Landing_Page/*", http.StripPrefix("/Static/Landing_Page/", http.FileServer(http.Dir("./Static/Landing_Page/"))))
	r.Handle("/Static/SpotifyPage/*", http.StripPrefix("/Static/SpotifyPage/", http.FileServer(http.Dir("./Static/SpotifyPage/"))))
	r.Handle("/Static/MediumPage/*", http.StripPrefix("/Static/MediumPage/", http.FileServer(http.Dir("./Static/MediumPage/"))))
	r.Handle("/Static/GalleryPage/*", http.StripPrefix("/Static/GalleryPage/", http.FileServer(http.Dir("./Static/GalleryPage/"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./Static/Landing_Page/test_me.html")
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		mydata := InputData{}
		err := json.NewDecoder(r.Body).Decode(&mydata)
		fmt.Printf("Requested Song: %s", mydata.Input)
		if err != nil {
			fmt.Printf("An Error occured: %v", err)
		}
		AddSong(mydata.Input)
	})

	http.ListenAndServe(":8080", r)
}
