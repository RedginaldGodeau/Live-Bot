package route

import (
	"backend/pkg/application"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(app *application.App) {
	app.Router.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "Hello")
	})
}
