package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var Addr = ":8080"

type Movie struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Poster      string    `json:"poster"`
	MovieUrl    string    `json:"movie_url"`
	IsPaid      bool      `json:"is_paid"`
	ReleaseYear time.Time `json:"release_year"`
	Genre       string    `json:"genre"`
}

func timeMustParse(year string) time.Time {
	t, err := time.Parse("2006", year)
	if err != nil {
		panic(err)
	}
	return t
}

func movieListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	mm := []Movie{
		Movie{1, "Бойцовский клуб", "/static/posters/fightclub.jpg", "https://www.youtube.com/watch?v=GbS-kM6jb9w", true, timeMustParse("1999"), "triller"},
		Movie{2, "Крестный отец", "/static/posters/father.jpg", "https://www.youtube.com/watch?v=otIzuKNaC0Qw", false, timeMustParse("1972"), "drama"},
		Movie{3, "Криминальное чтиво", "/static/posters/pulpfiction.jpg", "https://www.youtube.com/watch?v=iknlRtZ-A60", true, timeMustParse("1994"), "comedy"},
	}

	err := json.NewEncoder(w).Encode(mm)
	if err != nil {
		log.Printf("Render response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func main() {
	r := mux.NewRouter()
	r.ParseForm()
	r.HandleFunc("/movies", movieListHandler)
	http.Handle("/", r)
	log.Printf("Starting on port %s", Addr)
	log.Fatal(http.ListenAndServe(Addr, nil))
}
