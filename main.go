//go:generate air
//go:generate templ generate
//go:generate go run main.go
package main

import (
	"github.com/labstack/echo/v4"

	"github.com/jfsog/GoBlogado/handler"
)

func main() {
	e := echo.New()
	// e.Static("/static", "static")
	e.StaticFS(e.AcquireContext().Path(), e.Filesystem)
	uh := handler.UsersHandler{}
	e.GET("/", uh.HandleBase)
	e.Logger.Fatal(e.Start(":8080"))
}
