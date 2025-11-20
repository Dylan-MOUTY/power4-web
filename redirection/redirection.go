package redirection

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"

	"power-4/game"
)

var currentGame = game.NewGame(6, 7)

func RegisterRoutes() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/start", handleStart)
	http.HandleFunc("/game", handleGame)
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/reset", handleReset)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	openBrowser("http://localhost:7070/")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/Page1.html"))
	tmpl.Execute(w, nil)
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		switch r.FormValue("level") {
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

func handleGame(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, currentGame)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		col, err := strconv.Atoi(r.FormValue("column"))
		if err == nil {
			currentGame.PlayMove(col)
		}
	}
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	currentGame = game.NewGame(6, 7)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = exec.Command("xdg-open", url).Start()
	}
	if err != nil {
		fmt.Println("💡 Ouvre ton navigateur sur :", url)
	}
}
