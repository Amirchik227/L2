package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func anagram(words *[]string) *map[string]*[]string {
	// Мапа поиска анаграмм
	anagramSets := make(map[string]*[]string)

	// Приводим слова к нижнему регистру
	for k := range *words {
		(*words)[k] = strings.ToLower((*words)[k])
	}

	for _, word := range *words {
		// Сортируем буквы
		sortedWord := sortString(word)

		// Если множество анаграмм уже существует, добавляем слово в него
		if _, ok := anagramSets[sortedWord]; ok {
			*anagramSets[sortedWord] = append(*anagramSets[sortedWord], word)
		} else {
			// Создаем новое множество анаграмм и добавляем его в мапу
			anagramSets[sortedWord] = &[]string{word}
		}
	}

	// Мапа у которой ключём является первый элемент массива.
	// элементы массива отсортированы
	sortedAnagramSets := make(map[string]*[]string)

	for _, v := range anagramSets {
		// Множество из одного элемента не сохраяем
		if len(*v) == 1 {
			continue
		}
		// Сохраняем первое слово
		firstWord := (*v)[0]
		// Сортируем массив
		sort.Strings(*v)
		// Сохраняем сортированный массив
		sortedAnagramSets[firstWord] = v
	}

	return &sortedAnagramSets
}

// Функция соритровки букв в слове в алфавитном порядке
func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	dictionary := []string{"пятак", "пятка", "анаграммынету", "тяпка", "СТОлик", "листок", "слитОК", "клоун", "Кулон", "УКЛОН"}

	// Запускаем функцию
	result := anagram(&dictionary)

	// Вывод анаграмм
	for key, words := range *result {
		fmt.Printf("%v: %v\n", key, *words)
	}
}
