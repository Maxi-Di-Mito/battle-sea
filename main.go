package main

import (
	"fmt"
	"htmx-app/api/entities"
	"htmx-app/api/logic"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var CurrentGame *entities.Game

func main() {

	// db.InitDb()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HomeHandler)

	fmt.Println("RUNNING on 8081")
	p1 := logic.GetNewPlayer("Maxi")
	p2 := logic.GetNewPlayer("Cele")

	CurrentGame = logic.InitGame(p1, p2)

	err := http.ListenAndServe(":8765", nil)
	if err != nil {
		panic(err)
	}

}
