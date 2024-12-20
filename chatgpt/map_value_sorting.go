/*
В Go нельзя напрямую сортировать map, так как ключи в map хранятся в произвольном порядке.
Это одно из свойств map в Go: они оптимизированы для быстрого поиска, но не для хранения в отсортированном виде.

Чтобы "отсортировать" map, можно извлечь его ключи или значения в срез и отсортировать их.
Затем можно использовать отсортированные ключи для упорядоченного вывода элементов map.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Исходный map
	myMap := map[string]int{
		"apple":  5,
		"banana": 2,
		"cherry": 8,
		"date":   3,
	}

	// Создаем срез для хранения ключей
	keys := make([]string, 0, len(myMap))

	// Заполняем срез ключами
	for key := range myMap {
		keys = append(keys, key)
	}

	// Сортируем ключи по значениям
	sort.Slice(keys, func(i, j int) bool {
		return myMap[keys[i]] < myMap[keys[j]] // Меняем знак < на > для сортировки по убыванию
	})

	// Выводим отсортированные пары ключ-значение
	fmt.Println("Сортировка по значениям:")
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, myMap[key])
	}
}