package router

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouterInterface interface {
	Serve() error
}

type Router struct {
	echo.Echo
	Port string
}

func NewRouter(port string) *Router {
	r := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}
	r.Renderer = renderer

	return &Router{
		*r,
		port,
	}
}

func Secures(group *echo.Group, APIKey string) *echo.Group {
	group.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == APIKey, nil
		},
	}))

	return group
}

func (r *Router) Serve() error {

	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("FRONT_END_URL")},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
	}))

	r.Use(middleware.Logger())
	r.Use(middleware.Recover())
	r.Static("/upload", "upload")
	r.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public", http.FileServer(http.Dir("assets")))))

	err := r.Start(":" + r.Port)
	if err != nil {
		return err
	}

	return nil
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
