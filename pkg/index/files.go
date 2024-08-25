package index

import (
	"io"
	"log"
	"os"
)

// Функция удаляет файлы индекса для дальнейшей переиндексации сайтов и возвращает ошибку, если удалить не удалось
func DeleteIndexFile() (err error) {
	err = os.Remove(pagesFile)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln("error index file: ", err)
	}
	err = os.Remove(wordsFile)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln("error index file: ", err)
	}
	return err
}

// Функция откроет файлы и вернет(pages *os.File, words *os.File, err error)
func OpenIndexFile() (pages *os.File, words *os.File, reindexFlag bool, err error) {
	reindexFlag = false
	pages, err = os.Open(pagesFile)
	if err != nil {
		if os.IsNotExist(err) {
			pages, err = os.Create(pagesFile)
			if err != nil {
				log.Println(err)
			}
			reindexFlag = true
		} else {
			log.Println(err)
		}
	}
	words, err = os.Open(wordsFile)
	if err != nil {
		if os.IsNotExist(err) {
			words, err = os.Create(wordsFile)
			if err != nil {
				log.Println(err)
			}
			reindexFlag = true
		} else {
			log.Println(err)
		}
	}
	return pages, words, reindexFlag, nil
}

func (p *Pages) SaveIndexFile(pagesf *os.File, wordsf *os.File) (err error) {
	_, err = pagesf.Write(p.bytePages)
	if err != nil {
		log.Println("Save pages file error: ", err)
	}
	_, err = wordsf.Write(p.byteWords)
	if err != nil {
		log.Println("Save words file error: ", err)
	}
	return err
}

func (p *Pages) readfile(file *os.File, b *[]byte) (n int, err error) {
	buf := make([]byte, 100)
	for {
		n, err = file.Read(buf)
		if err != nil {
			if err == io.EOF {
				return n, nil
			}
			return n, err
		}
		*b = append(*b, buf[:n]...)
	}
	// буферизиванное чтение
}
