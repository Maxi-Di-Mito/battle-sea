package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "index", CurrentGame)
}

func ClickCell(w http.ResponseWriter, r *http.Request) {

}
