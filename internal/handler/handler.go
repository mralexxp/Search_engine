package handler

import (
	"bufio"
	"fmt"
	"gosearch/pkg/crawler"
	"gosearch/pkg/index"
	"log"
	"net"
)

func responser(b []byte, conn net.Conn) (err error) {
	b = append(b, []byte("\n\r")...)
	_, err = conn.Write(b)
	return err
}

func Handler(conn net.Conn, sch *index.Pages) {
	defer conn.Close()

	rd := bufio.NewReader(conn)

	for {
		msg, _, err := rd.ReadLine()
		if err != nil {
			if _, ok := err.(*net.OpError); ok {
				fmt.Println("Connection closed")
				return
			} else {
				log.Fatal(err)
			}
		}
		res := sch.Search(string(msg))

		// Установить дедлайн на чтение клиентом
		// должны вернуть []crawler.Document
		err = responser(crawler.DocumentSerialize(&res), conn)
		if err != nil {
			if _, ok := err.(*net.OpError); ok {
				fmt.Println("Connection closed")
				return
			} else {
				log.Fatal(err)
			}
		}
	}
}
