package main

import (
	"bufio"
	"fmt"
	"gosearch/pkg/crawler"
	"log"
	"net"
)

const (
	targetIP = "127.0.0.1:12345"
)

func printResult(data []crawler.Document, target string) {
	if len(data) == 0 {
		fmt.Printf("\"%v\" not found. Please, edit search string or edit depth\n", target)
		return
	}
	fmt.Printf("Found %d documents on the \"%v\" query:\n", len(data), target)
	for i, value := range data {
		fmt.Printf("%d: %v\nURL: %v\n=============================\n", i+1, value.Title, value.URL)
	}
}

func reader(rd bufio.Reader) (b []byte) {
	for {
		buf, _, err := rd.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		b = append(b, buf...)
		if len(buf) < 4096 {
			break
		}
	}
	return b
}

// Дедлайны соединения
func main() {
	fmt.Println("Connecting to", targetIP)
	conn, err := net.Dial("tcp4", targetIP)
	if err != nil {
		log.Fatalf("Error connecting to %v - %v", targetIP, err)
	}
	rd := bufio.NewReader(conn)
	for {
		var s string
		fmt.Print("Enter search text: ")
		fmt.Scan(&s)
		_, err := conn.Write([]byte(s + "\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Ожидаем ответа...")
		msg := reader(*rd)
		printResult(crawler.DocumentDeSerialize(msg), s)
	}
}
