package main

import (
	"html/template"
	"htmx-app/api/entities"
	"htmx-app/api/logic"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var tmpFile = "index.go.html"

	tmpl, _ := template.New(tmpFile).ParseFiles(tmpFile)

	w.Header().Add("Content-Type", "text/html")
	tmpl.Execute(w, CurrentGame)

}

func StartGame(w http.ResponseWriter, r *http.Request) *entities.Game {
	p1 := logic.GetNewPlayer("Maxi")
	p2 := logic.GetNewPlayer("Cele")

	game := logic.InitGame(p1, p2)

	return game
}
