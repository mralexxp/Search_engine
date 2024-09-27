package webapp

import (
	"gosearch/pkg/crawler"
	"gosearch/pkg/index"
	"net/http"

	"github.com/gorilla/mux"
)

type WebApp struct {
	addr  string
	Pages []crawler.Document
	Words map[string][]int
}

// Constructor
func NewWebApp(addr string, p *index.Pages) (w *WebApp) {
	pages := p.GetPages()
	words := p.GetWords()
	return &WebApp{
		addr:  addr,
		Pages: pages,
		Words: words,
	}
}

func (w *WebApp) Start() (err error) {
	m := mux.NewRouter()

	w.endpoints(m)

	err = http.ListenAndServe(w.addr, m)
	if err != nil {
		return err
	}
	return err
}

func (w *WebApp) endpoints(m *mux.Router) {
	m.HandleFunc("/pages", w.pagesHandler)
	m.HandleFunc("/words", w.wordsHandler)
}
