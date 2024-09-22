package main

import (
	"fmt"
	"gosearch/internal/handler"
	"gosearch/pkg/index"
	"log"
	"net"
)

const (
	IndexPagesfile = "pages.bin" // Файл индексации страниц
	IndexWordsFile = "words.bin" // Файл индексации ключевых слов
	depth          = 2           // Глубина сканирования сайтов
	reindex        = false
	ListenAdress   = "127.0.0.1:12345"
)

// func getFlagS() (string, bool) {
// 	reindex := true
// 	target := flag.String("s", "", "Enter string for search in sites")
// 	index := flag.String("i", "", "Use \"-i on\" for reindex sites")
// 	flag.Parse()
// 	if *index == "" {
// 		reindex = false
// 	}
// 	return *target, reindex
// }

// later: go scan
func main() {
	if reindex {
		fmt.Println("Indexing on. Starting service...")
		err := index.DeleteIndexFile()
		if err != nil {
			log.Println(err)
		}
	}

	urls := []string{"https://go.dev/"}

	search := index.New(urls, depth)

	fmt.Println("Server starting...")
	listener, err := net.Listen("tcp4", ListenAdress)
	if err != nil {
		log.Fatalf("Error listening to %v - %v", ListenAdress, err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accept connection to %v - %v", conn.RemoteAddr().String(), err)
		}
		fmt.Println("Connection established. Waiting command...")
		go handler.Handler(conn, search)
	}

	// printResult(search.Search(target), target)

}
