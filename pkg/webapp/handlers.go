package webapp

import (
	"fmt"
	"net/http"
)

func (wa *WebApp) pagesHandler(w http.ResponseWriter, r *http.Request) {
	pages := wa.Pages
	for i, document := range pages {
		fmt.Fprintf(w, "======[%d]=======\nDocumentID: %d\nTitle: %v\nURL: %v\n", i+1, document.ID, document.Title, document.URL)

	}
	fmt.Fprintf(w, "======[END]=======")
}

func (wa *WebApp) wordsHandler(w http.ResponseWriter, r *http.Request) {
	words := wa.Words
	i := 0
	for s, document := range words {
		i++
		fmt.Fprintf(w, "%d: %s[%d]\n", i, s, len(document))
		for _, id := range document {
			fmt.Fprintf(w, "    ID: %v\n", id)
		}

	}
	fmt.Fprintf(w, "======[END]=======")
}
