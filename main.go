package main

import (
	"fmt"
	"html/template"
	"htmx-app/api/logic"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Person struct {
	Name string
	Age  int
}

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

	server := echo.New()
	server.Use(middleware.Logger())
	server.Static("/static", "static")
	server.Renderer = temps

	server.GET("/:player", HomeHandler)
	server.POST("/click-cell", ClickCellHandler)

	p1 := logic.GetNewPlayer("Maxi")
	p2 := logic.GetNewPlayer("Cele")

	logic.CurrentGame = logic.InitGame(p1, p2)

	server.Logger.Fatal(server.Start(":8765"))
	fmt.Println("RUNNING on 8765")
}
