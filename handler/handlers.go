package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/jfsog/GoBlogado/view/layout"
	"github.com/jfsog/GoBlogado/view/user"
)

type UsersHandler struct{}

func (h UsersHandler) HandleUserShow(c echo.Context) error {
	return render(c, user.Show())
}

func (h UsersHandler) HandleBase(c echo.Context) error {
	return render(c, layout.Base(layout.Posts()))
}

func (h UsersHandler) HandleAbout(c echo.Context) error {
	return render(c, layout.About())
}

func (h UsersHandler) HandlePosts(c echo.Context) error {
	return render(c, layout.Posts())
}
