package main

import (
	"fmt"
	"net/http"

	"github.com/RenanHCosta/tarta/templates"
	"github.com/RenanHCosta/tarta/templates/pages"
)

const WEBSITE_TITLE string = "Tarta"

func isHtmxDrivenRequest(r *http.Request) bool {
	return r.Header.Get("hx-request") == "true"
}

func main() {
	port := ":8080"

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/*", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if isHtmxDrivenRequest(r) {
			pages.Home().Render(r.Context(), w)
		} else {
			templates.Layout(pages.Home(), WEBSITE_TITLE).Render(r.Context(), w)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		if isHtmxDrivenRequest(r) {
			pages.About().Render(r.Context(), w)
		} else {
			templates.Layout(pages.About(), WEBSITE_TITLE).Render(r.Context(), w)
		}
	})

	fmt.Println("Server started on port", port)
	http.ListenAndServe(port, nil)
}
