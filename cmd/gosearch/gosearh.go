package main

import (
	"flag"
	"fmt"
	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"strings"
	// "gosearch/pkg/crawler/membot"
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

func search(text string, res *[]crawler.Document) (result []crawler.Document, qty int) {
	for _, page := range *res {
		contain := strings.Contains(page.Title, text)
		if contain {
			result = append(result, page)
			fmt.Println(page.URL)
			fmt.Println(page.Title)
			qty++
		}
	}
	return result, qty

}

func main() {
	urls := []string{"https://go.dev"}
	depth := 2
	sText := getFlagS()
	res := craw(urls, depth)
	for _, value := range res {
		fmt.Println("========================")
		fmt.Println(value.ID)
		fmt.Println(value.URL)
		fmt.Println(value.Title)
		fmt.Println(value.Body)
		fmt.Println("========================")
	}
	if sText == "" {
		fmt.Println("Аргументы отсутствуют. Поиск отменен.")
	} else {
		fmt.Println("Поиск по запросу: ", sText)
		_, qty := search(sText, &res) // Возвращает слайс из crawler.Document
		fmt.Println("Найдено (кол-во): ", qty)
	}

}
