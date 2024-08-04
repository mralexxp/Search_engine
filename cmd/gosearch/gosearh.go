package main

import (
	"flag"
	"fmt"
	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"gosearch/pkg/index"
)

func getFlagS() string {
	param := flag.String("s", "", "Enter string for search in sites")
	flag.Parse()
	return *param
}

func craw(urls []string, depth int) []crawler.Document {
	var reSearch = []crawler.Document{}
	for _, url := range urls {
		craw := spider.New()
		result, err := craw.Scan(url, depth)
		if err != nil {
			panic("Обшибочка вышел")
		}
		reSearch = append(reSearch, result...)
	}
	return reSearch
}

func main() {
	urls := []string{"https://go.dev/"}
	depth := 2
	sText := getFlagS()
	// sText := "Packages"
	// Временное решение с ветвлением VVV
	switch sText {
	case "":
		fmt.Println("Аргументы отсутствуют. Поиск отменен.")
	default:
		fmt.Println("Поиск по запросу: ", sText)
		allDocuments := craw(urls, depth)
		pages := index.New(allDocuments)
		tempTest := pages.Search(sText)
		if len(tempTest) != 0 {
			fmt.Println("Найденные документы по запросу \"", sText, "\":")
			for _, value := range tempTest {
				fmt.Printf("ID: %d \nTitle: %v \nURL: %v \n=========================\n", value.ID, value.Title, value.URL)
			}
		} else {
			fmt.Printf("По запросу \"%v\" - не найдено ни одного документа", sText)
		}
	}

}
