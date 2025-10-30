package main

import (
	"html/template"
	"net/http"

	"power-4/game"
)

var currentGame *game.Game

func main() {
	// Routes
	http.HandleFunc("/", handlePage1)      // page d'accueil : choix du niveau
	http.HandleFunc("/start", handleStart) // création partie selon le niveau
	http.HandleFunc("/game", handleIndex)  // affichage de la grille
	http.HandleFunc("/play", handlePlay)   // jouer un coup
	http.HandleFunc("/reset", handleReset) // recommencer la partie

	// Démarrer le serveur
	http.ListenAndServe(":8080", nil)
}

// Page d'accueil (Page1.html)
func handlePage1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/Page1.html"))
	tmpl.Execute(w, nil)
}

// Démarrage d'une nouvelle partie selon le niveau choisi
func handleStart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		level := r.FormValue("level")
		switch level {
		case "easy":
			currentGame = game.NewGame(6, 7)
		case "normal":
			currentGame = game.NewGame(6, 9)
		case "hard":
			currentGame = game.NewGame(7, 8)
		default:
			currentGame = game.NewGame(6, 7)
		}
	}
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

// Page principale du jeu (index.html)
func handleIndex(w http.ResponseWriter, r *http.Request) {
	if currentGame == nil {
		currentGame = game.NewGame(6, 7) // par défaut
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, currentGame)
}

// Quand le joueur joue un coup
func handlePlay(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && currentGame != nil {
		col := r.FormValue("column")
		column := 0
		for _, c := range col {
			column = int(c - '0')
		}
		currentGame.PlayMove(column)
	}
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

// Réinitialisation de la partie (même niveau qu'avant)
func handleReset(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if currentGame != nil {
			rows := len(currentGame.Board)
			cols := len(currentGame.Board[0])
			currentGame = game.NewGame(rows, cols)
		} else {
			currentGame = game.NewGame(6, 7)
		}
	}
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}
