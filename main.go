//go:generate air
//go:generate templ generate
//go:generate go run main.go
package main

import (
	"os"

	"github.com/labstack/echo/v4"

	"github.com/jfsog/GoBlogado/handler"
)

func main() {
	e := echo.New()
	e.StaticFS(e.AcquireContext().Path(), e.Filesystem)
	uh := handler.UsersHandler{}
	e.GET("/", uh.HandleBase)
	e.GET("/about", uh.HandleAbout)
	e.GET("/posts", uh.HandlePosts)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
