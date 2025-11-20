package main

import (
	"fmt"
	"log"
	"net/http"

	"power-4/handlers"
)

func main() {
	handlers.RegisterRoutes()
	fmt.Println("🚀 Serveur lancé sur http://localhost:7070/")
	log.Fatal(http.ListenAndServe(":7070", nil))
}
