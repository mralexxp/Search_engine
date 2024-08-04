package index

import (
	"gosearch/pkg/crawler"
	"strings"
)

// Индекс: map[слово]массив номеров документов, где встречается это слово
// Функция принимает структуры Document, разбивает на слова и возвращает map[слово][]ID-документов, где это слово встречается
func IndexSplitter(doc []crawler.Document) map[string][]int {
	result := make(map[string][]int)
	for _, value := range doc {
		for _, word := range strings.Split(value.Title, " ") {
			idx := result[word]
			idx = append(idx, value.ID)
			result[word] = idx
		}
	}
	return result
}
