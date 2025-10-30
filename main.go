package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"

	"power-4/game"
)

var currentGame *game.Game

func main() {
	currentGame = game.NewGame() // par défaut : easy

	// Routes
	http.HandleFunc("/", handleIndex)      // Page d'accueil (Page1.html)
	http.HandleFunc("/start", handleStart) // Création d'une partie selon le niveau
	http.HandleFunc("/game", handleGame)   // Page du jeu (index.html)
	http.HandleFunc("/play", handlePlay)   // Jouer un coup
	http.HandleFunc("/reset", handleReset) // Recommencer
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	url := "http://localhost:8080/"
	fmt.Println("🚀 Serveur lancé sur :", url)
	openBrowser(url) // ouvre automatiquement ton navigateur

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// --------------------- HANDLERS ---------------------

// Page d'accueil — choix du niveau
func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/Page1.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Démarre une partie selon le niveau choisi
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

// Page principale du jeu
func handleGame(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, currentGame); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Quand le joueur clique dans une colonne
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

// Bouton "Rejouer"
func handleReset(w http.ResponseWriter, r *http.Request) {
	currentGame = game.NewGame()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// --------------------- UTILITAIRE ---------------------

// Ouvre automatiquement le navigateur par défaut
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default: // linux
		err = exec.Command("xdg-open", url).Start()
	}
	if err != nil {
		fmt.Println("💡 Ouvre ton navigateur sur :", url)
	}
}
