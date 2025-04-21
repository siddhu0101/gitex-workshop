package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {
	e := echo.New()

	// Serve static files
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", http.FileServer(http.FS(staticFS)))))

	// Serve the index page
	e.GET("/", func(c echo.Context) error {
		tmpl := template.Must(template.ParseFS(templatesFS, "templates/index.html"))
		data := struct{ Title, Message string }{
			Title:   "Gitex Asia 2025: GitOps with ArgoCD",
			Message: "Welcome to our hands-on workshop! Deployed via ArgoCD on KodeKloud.",
		}
		return tmpl.Execute(c.Response().Writer, data)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
