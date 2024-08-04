package index

import (
	"gosearch/pkg/crawler"
	generators "gosearch/pkg/utils"
	"sort"
)

// Кэш документов должен состоять из слайса: структуры документов, сортировать которую будем методом обращения к ID-документа.
// Бинарный поиск необходимого ID-документа мы будем также отсуществлять с помощью обращенийк ID в структуре.

type pages struct {
	pages []crawler.Document
	words map[string][]int
}

func New(allPages []crawler.Document) *pages {
	return &pages{
		pages: sortPages(allPages), // Сортированные по ID
		words: IndexSplitter(allPages),
	}
}

func sortPages(documents []crawler.Document) []crawler.Document {
	sort.Slice(documents, func(i, j int) bool {
		return documents[i].ID < documents[j].ID
	})
	return documents
}

// Функция должна вернуть структуры из документов, в которых встречаются слова TARGET
func (p *pages) Search(target string) []crawler.Document {
	result := make([]crawler.Document, 0)
	exist := make([]int, 0)
	for _, ids := range p.words[target] {
		indx := generators.SimpleSearch(exist, ids)
		if indx == -1 {
			exist = append(exist, ids)
			result = append(result, p.pages[generators.BinarySearch(p.pages, ids)])
		}
	}
	return result
}
