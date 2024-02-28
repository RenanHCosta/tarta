package main

import (
	"fmt"
	"net/http"

	"github.com/RenanHCosta/tarta/integrations/vtex"
	"github.com/RenanHCosta/tarta/templates"
	"github.com/RenanHCosta/tarta/templates/components"
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

	input := vtex.ProductSearchInput{
		From:                 0,
		To:                   9,
		FullText:             "",
		OrderBy:              "OrderByScoreDESC",
		PriceRange:           "",
		SalesChannel:         "",
		HideUnavailableItems: false,
		SimulationBehavior:   vtex.SimulationBehaviorDefault,
		SelectedFacets:       []vtex.SelectedFacetInput{},
	}

	http.HandleFunc("/product-search", func(w http.ResponseWriter, r *http.Request) {
		components.RenderProductCards(vtex.ProductSearch(input).ProductSearch.Products).Render(r.Context(), w)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			templates.Layout(templates.NotFound()).Render(r.Context(), w)
			return
		}

		if isHtmxDrivenRequest(r) {
			pages.Home(input).Render(r.Context(), w)
		} else {
			templates.Layout(pages.Home(input)).Render(r.Context(), w)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		if isHtmxDrivenRequest(r) {
			pages.About().Render(r.Context(), w)
		} else {
			templates.Layout(pages.About()).Render(r.Context(), w)
		}
	})

	fmt.Println("Server started on port", port)
	http.ListenAndServe(port, nil)
}
