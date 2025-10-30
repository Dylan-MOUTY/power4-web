package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"power-4/game"
)

var currentGame *game.Game

func main() {
	currentGame = game.NewGame()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/game", handleGame)
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/reset", handleReset)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/Page1.html"))
	if err := tmpl.Execute(w, currentGame); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleGame(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, currentGame); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		colStr := r.FormValue("column")
		col, err := strconv.Atoi(colStr)
		if err == nil {
			currentGame.PlayMove(col)
		}
	}
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	currentGame = game.NewGame()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
