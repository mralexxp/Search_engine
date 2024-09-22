package crawler

import (
	"encoding/json"
	"log"
)

// Поисковый робот.
// Осуществляет сканирование сайтов.

// Interface определяет контракт поискового робота.
type Interface interface {
	Scan(url string, depth int) ([]Document, error)
	BatchScan(urls []string, depth int, workers int) (<-chan Document, <-chan error)
}

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int    `json:"ID"`
	URL   string `json:"URL"`
	Title string `json:"Title"`
}

func DocumentSerialize(doc *[]Document) (b []byte) {
	result, err := json.Marshal(*doc)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func DocumentDeSerialize(b []byte) (doc []Document) {
	err := json.Unmarshal(b, &doc)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
