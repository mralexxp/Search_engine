package generators

import (
	"gosearch/pkg/crawler"
)

func SimpleSearch(data []int, item int) (i int) {
	for i := range data {
		if data[i] == item {
			return i
		}
	}
	return -1
}

// Функция должна принять слайс из документов
func BinarySearch(data []crawler.Document, id int) int {
	low, high := 0, len(data)-1 // Минимум и максимум по умолчанию 0 и длина минус 1
	for low <= high {           // Цикл, пока минимум не будет больше либо равен максимуму
		mid := (low + high) / 2 // Середина = минимум + максимум / 2
		if data[mid].ID == id { // Если нашли значение в data[середина]
			return mid // Возвращаем индекс массива с найденным значением
		}
		if data[mid].ID < id { // Если значение середины меньше искомого
			low = mid + 1 // минимумом становится середина + 1
		} else { // Если значение середины больше искомого
			high = mid - 1 // максимум становится серединой - 1
		}
	}
	return -1 // Возвращаем -1, если значение не найдено.
}
