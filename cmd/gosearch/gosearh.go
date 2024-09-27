package main

import (
	"fmt"
	"gosearch/internal/handler"
	"gosearch/pkg/index"
	"gosearch/pkg/webapp"
	"log"
	"net"
)

const (
	IndexPagesfile = "pages.bin" // Файл индексации страниц
	IndexWordsFile = "words.bin" // Файл индексации ключевых слов
	depth          = 2           // Глубина сканирования сайтов
	reindex        = false
	ListenAdress   = "127.0.0.1:12345"
	WebAppAddress  = "127.0.0.1:8080"
)

// Дальнешйее развитие многоканального парсера требует переработки пакета spider

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

	wa := webapp.NewWebApp(WebAppAddress, search)
	err := wa.Start()
	if err != nil {
		log.Fatal(err)
	}

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
}
