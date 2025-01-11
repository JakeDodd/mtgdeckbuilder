package main

import (
	"html/template"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Message struct {
	Message string
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}
	port = ":" + port

	// e is an Echo http server, see above import
	// dont know much about this library
	e := echo.New()
	e.Use(middleware.Logger())

	// TODO: return to this later once database is seeded.
	//e.GET("/rand-card", func(c echo.Context) error {
	//	return c.Render(200, "index", message)
	//})

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper:    nil,
		Root:       "react/frontend/dist",
		Index:      "index.html",
		HTML5:      true,
		Browse:     false,
		IgnoreBase: false,
		Filesystem: nil,
	}))

	log.Printf("Server starting on port: %v", port)

	// Here we start the server, this starts a loop which will only end if an error
	// is thrown while processing a request, which is why we log.Fatal the returned value
	e.Logger.Fatal(e.Start(port))

}
