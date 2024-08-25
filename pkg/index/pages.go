package index

import (
	"encoding/json"
	"gosearch/pkg/crawler"
	generators "gosearch/pkg/utils"
	"log"
	"sort"
)

// Кэш документов должен состоять из слайса: структуры документов, сортировать которую будем методом обращения к ID-документа.
// Бинарный поиск необходимого ID-документа мы будем также отсуществлять с помощью обращенийк ID в структуре.

// type Pages interface {
// 	Search(string) []crawler.Document
// 	New([]crawler.Document) *pages
// }

type Pages struct {
	pages     []crawler.Document
	words     map[string][]int
	bytePages []byte
	byteWords []byte
}

func (p *Pages) GetInfo() (pages *[]crawler.Document, words *map[string][]int) {
	return &p.pages, &p.words
}

// Конструктор должен читать из файла индекс, собирать структуру и проверять наличие записей
func New(urls []string, depth int) (search *Pages) {
	search = &Pages{
		pages: []crawler.Document{},
		words: map[string][]int{},
	}

	pagesf, wordsf, reindexFlag, err := OpenIndexFile()
	defer func() {
		err = pagesf.Close()
		if err != nil {
			log.Println("Error close pagefile", err)
		}
		err = wordsf.Close()
		if err != nil {
			log.Println("Error close wordsfile", err)
		}
	}()
	if err != nil {
		log.Println("Error open index file: ", err)
	}
	// 		2.1 Если нет файлов: мы создали новые и начинаем собирать структуру
	if reindexFlag {
		// 		2.1.1 Сканируем сайты и формируем []crawler.Document
		search.pages, err = craw(urls, depth)
		if err != nil {
			log.Println("Crawler error: ", err)
		}
		// 		2.1.1.1 Сортируем []crawler.Document (с помощью sortPages())
		search.pages = sortPages(search.pages)
		// 		2.1.2 Формируем words map[string][]int (Indexer)
		err := search.Indexer()
		if err != nil {
			log.Println("Indexer error: ", err)
		}
		// 		2.1.3 Сериализуем (struct => json => []byte)
		err = search.SerializePages()
		if err != nil {
			log.Println("Serialize pages error: ", err)
		}
		err = search.SerializeWords()
		if err != nil {
			log.Println("Serialize words error: ", err)
		}
		// 		2.1.4 Сохраняем индекс pages и words в файлы
		err = search.SaveIndexFile(pagesf, wordsf)
		if err != nil {
			log.Println(err)
		}
	} else {
		// 			2.2.1 Загружаем файлы
		// Читаем pagesfile:
		_, err := search.readfile(pagesf, &search.bytePages)
		if err != nil {
			log.Println("Error read pagesfile: ", err)
		}
		// Читаем wordsfile
		_, err = search.readfile(wordsf, &search.byteWords)
		if err != nil {
			log.Println("Error read pagesfile: ", err)
		}
		// ============================================
		// 			2.2.2 Десериализуем ([]byte => JSON => struct)
		_, err = search.DeSerializePages(&search.pages)
		if err != nil {
			log.Println(err)
		}
		_, err = search.DeSerializeWords(&search.words)
		if err != nil {
			log.Println(err)
		}
	}

	// 		2.2 Если файлы есть, то:

	// 			2.2.1 Загружаем файлы
	// 			2.2.2 Десериализуем ([]byte => JSON => struct)
	// 			2.2.3 Собираем структуру
	search.bytePages, search.byteWords = []byte{}, []byte{}
	return search

	// 	return &Pages{
	//		pages: sortPages(allPages), // Сортированные по ID
	// 		words: Indexer(allPages),
	// }
}

// Сортирует документы по ID-документов в порядке возрастания
// TODO:
// - попробовать принимать указатель на слайс и возвращать его же
func sortPages(documents []crawler.Document) []crawler.Document {
	sort.Slice(documents, func(i, j int) bool {
		return documents[i].ID < documents[j].ID
	})
	return documents
}

// Функция должна вернуть структуры из документов, в которых встречаются слова TARGET
func (p *Pages) Search(target string) []crawler.Document {
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

// Четыре метода ниже должны реализовать интерфейс для четния и записи в файл.

func (pages *Pages) SerializePages() (err error) { // Выход заменить на интерфейс io.Writer??
	pages.bytePages, err = json.MarshalIndent(pages.pages, "", "  ")
	if err != nil {
		return err
	}
	return err
}

func (pages *Pages) DeSerializePages(pagesStruct *[]crawler.Document) (n int, err error) {
	err = json.Unmarshal(pages.bytePages, pagesStruct)
	if err != nil {
		log.Println("Error unmarshal JSON pagesbyte: ", err)
	}
	return 0, err
} // вход заменить на io.Reader

func (pages *Pages) SerializeWords() (err error) {
	pages.byteWords, err = json.MarshalIndent(pages.words, "", "  ")
	if err != nil {
		return err
	}
	return err
}

func (pages *Pages) DeSerializeWords(wordsStruct *map[string][]int) (n int, err error) {
	err = json.Unmarshal(pages.byteWords, wordsStruct)
	if err != nil {
		log.Println("Error unmarshal JSON wordsbyte: ", err)
	}
	return 0, err
} // вход заменить на io.Reader
