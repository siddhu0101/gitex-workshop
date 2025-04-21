package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed templates/*
var tmplFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {
	addr := getEnv("PORT", "8080")

	// parse templates at startup (panic if invalid)
	tmpl := template.Must(template.ParseFS(tmplFS, "templates/index.html"))

	// serve static assets
	http.Handle("/static/", http.FileServer(http.FS(staticFS)))

	// main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct{ Title string }{
			Title: "Welcome to the Gitex Asia Workshop — Multi‑Tenancy Powered by Loft Labs",
		}
		if err := tmpl.Execute(w, data); err != nil {
			log.Printf("template execute error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	log.Printf("Starting server on %s…", addr)
	if err := http.ListenAndServe(":"+addr, nil); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
