package main

import (
	"fmt"
	"html/template"
	"htmx-app/api/entities"
	"htmx-app/api/logic"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Person struct {
	Name string
	Age  int
}

var CurrentGame *entities.Game

type Template struct {
	templates *template.Template
}

var temps = &Template{
	templates: template.Must(template.ParseGlob("*.go.html")),
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	// db.InitDb()
	fs := http.FileServer(http.Dir("./static"))

	server := echo.New()
	server.Use(middleware.Logger())
	server.Static("/static", "static")
	server.Renderer = temps
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	server.GET("/", HomeHandler)

	p1 := logic.GetNewPlayer("Maxi")
	p2 := logic.GetNewPlayer("Cele")

	CurrentGame = logic.InitGame(p1, p2)

	server.Logger.Fatal(server.Start(":8765"))
	fmt.Println("RUNNING on 8765")
}
