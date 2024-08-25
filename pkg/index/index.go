package index

import (
	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"log"
	"strings"
)

const (
	pagesFile = "pages.bin"
	wordsFile = "words.bin"
)

// Индекс: map[слово]массив номеров документов, где встречается это слово
// Функция принимает структуры Document, разбивает на слова и возвращает map[слово][]ID-документов, где это слово встречается
// func Indexer(doc []crawler.Document) map[string][]int {
func (p *Pages) Indexer() (err error) {
	for _, value := range p.pages {
		for _, word := range strings.Split(value.Title, " ") {
			idx := p.words[word]
			idx = append(idx, value.ID)
			p.words[word] = idx
		}
	}
	return err
}

// Метод pages принимает URLs и глубину, сканирует сайт и формирует []crawler.Document
func craw(urls []string, depth int) (docs []crawler.Document, err error) {
	// var reSearch = []crawler.Document{}
	for _, url := range urls {
		craw := spider.New()
		result, err := craw.Scan(url, depth)
		if err != nil {
			log.Println(err)
		}
		docs = append(docs, result...)
	}
	return docs, err
}
