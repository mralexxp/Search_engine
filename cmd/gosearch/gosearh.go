package main

import (
	"flag"
	"fmt"
	"gosearch/pkg/crawler"
	"gosearch/pkg/index"
	"log"
	"os"
)

const (
	IndexPagesfile = "pages.bin" // Файл индексации страниц
	IndexWordsFile = "words.bin" // Файл индексации ключевых слов
	depth          = 2           // Глубина сканирования сайтов
)

// Собирает аргументы запуска программы. Возвращает строку для поиска и reindex bool
func getFlagS() (string, bool) {
	reindex := true
	target := flag.String("s", "", "Enter string for search in sites")
	index := flag.String("i", "", "Use \"-i on\" for reindex sites")
	flag.Parse()
	if *index == "" {
		reindex = false
	}
	return *target, reindex
}

func printResult(data []crawler.Document, target string) {
	if len(data) == 0 {
		fmt.Printf("\"%v\" not found. Please, edit search string or edit depth", target)
		return
	}
	fmt.Printf("Found %d documents on the \"%v\" query:\n", len(data), target)
	for i, value := range data {
		fmt.Printf("%d: %v\nURL: %v\n=============================\n", i+1, value.Title, value.URL)
	}

}

// Алгоритм программы:
// Читаем флаги
// +1 Если есть флажок index, то:
// 		+1.1 Удаляем файлы индекса
//		+1.2 Триггер на новый индекс
// 2 Создаем объект поиска с помощью конструктора New()
// 		2.1 Если нет файлов, то:
// 			2.1.1 Сканируем сайты и формируем []crawler.Documents
// 			2.1.2 Формируем words map[string][]int
// 			2.1.3 Сериализуем (struct => json => []byte)
// 			2.1.4 Сохраняем индекс pages и words в файлы
// 		2.2 Если файлы есть, то:
// 			2.2.1 Загружаем файлы
// 			2.2.2 Десериализуем ([]byte => JSON => struct)
// 3 Если поисковый запрос пустой:
// 		3.1 Завершаем работу
// 4 Ищем поисковое слово в структуре в words, возвращаем []int, которй содержит id-документов
// 5 Ищем бинарным поиском все crawler.Document из []int, возвращаем []crawler.Document
// 6 Выводим результат: "Title: %v\nURL: %v\nID: %v", title, url, id

func main() {
	// Читаем флаги запуска программы
	target, reindex := getFlagS()
	// Если флаг -i содержит символы
	if reindex {
		fmt.Println("Indexing on. Started...")
		err := index.DeleteIndexFile()
		if err != nil {
			log.Println(err)
		}
	}
	urls := []string{"https://go.dev/"}
	search := index.New(urls, depth)
	// 3 Если поисковый запрос пустой: 3.1 Завершаем работу
	if target == "" {
		fmt.Println("Search string is empty. Exiting...")
		os.Exit(0)
	}
	printResult(search.Search(target), target)

}
