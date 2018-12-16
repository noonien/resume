package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/noonien/resume/resume"
	"github.com/noonien/resume/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gobuffalo/packr"

)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.Recoverer)

	var staticBox = packr.NewBox("static")

	// register static routes
	r.Get("/", resumeHandler)
	r.Get("/*", serveFiles(staticBox))

	err := http.ListenAndServe(":"+port, r)
	log.Fatal(err)

}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get("User-Agent"), "curl") {
		curlHandler(w, r)
		return
	}

	tmpl, err := template.HTML()
	if err != nil {
		log.Print(err)
		return
	}

	err = tmpl.Execute(w, resume.Resume())
	if err != nil {
		log.Print(err)
	}
}

func serveFiles(box packr.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f, err := box.Open(r.URL.Path)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		f.Close()

		fi, err := f.Stat()
		if err != nil || fi.IsDir() {
			http.NotFound(w, r)
			return
		}

		http.ServeContent(w, r, r.URL.Path, fi.ModTime(), f)
	}
}
