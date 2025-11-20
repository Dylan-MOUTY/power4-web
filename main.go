package main

import (
	"fmt"
	"log"
	"net/http"

	"power-4/redirection"
)

func main() {
	redirection.RegisterRoutes()
	fmt.Println("🚀 Serveur lancé sur http://localhost:7070/")
	log.Fatal(http.ListenAndServe(":7070", nil))
}
