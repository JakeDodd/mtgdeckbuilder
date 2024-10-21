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
	e.Renderer = newTemplate()
	var message Message = Message{Message: "Hello World"}

	// Here we define a GET endpoint at /rand-card that returns index.html
	// and passes the Message struct variable
	e.GET("/rand-card", func(c echo.Context) error {
		return c.Render(200, "index", message)
	})

	log.Printf("Server starting on port: %v", port)

	// Here we start the server, this starts a loop which will only end if an error
	// is thrown while processing a request, which is why we log.Fatal the returned value
	e.Logger.Fatal(e.Start(port))

}
