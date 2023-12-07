package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Значения ключей
var (
	fColumn  int
	fNumeric bool
	fReverse bool
	fUnique  bool
	fMonths  bool
	fSpaces  bool
	fSort    bool
	isSorted bool = true
)

// Номера месяцев для соритровки по месяцам
var month = map[string]byte{
	"jan": 1,
	"feb": 2,
	"mar": 3,
	"apr": 4,
	"may": 5,
	"jun": 6,
	"jul": 7,
	"aug": 8,
	"sep": 9,
	"oct": 10,
	"nov": 11,
	"dec": 12,
}

// Считывание ключей при запуске программы
func init() {
	flag.IntVar(&fColumn, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&fNumeric, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&fReverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&fUnique, "u", false, "не выводить повторяющиеся строки")

	flag.BoolVar(&fMonths, "M", false, "сортировка по месяцу")
	flag.BoolVar(&fSpaces, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&fSort, "c", false, "проверять отсортированы ли данные")
}

func main() {
	// Считывание флагов
	flag.Parse()

	// Строки для сортировки
	lines := []string{
		"b3g 29   dec",
		"  a2e 931  aug",
		"c4g 742  jun",
		"d7s 75   nov",
		" f1g 133  may",
	}

	// Запускаем функцию
	lines = mySort(lines)

	// Флаг - проверка отсортированы ли данные
	if isSorted {
		fmt.Println("Данные отсортированы")
	} else {
		fmt.Println("Данные не отсортированы")
	}

	// Вывод отсортированных строк
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Удаление повторяющихся срок
func removeDuplicates(lines []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			result = append(result, line)
		}
	}

	return result
}

func mySort(lines []string) []string {
	sort.SliceStable(lines, func(i, j int) bool {

		// Флаг игнорирования пробелов
		if fSpaces {
			lines[i] = strings.TrimSpace(lines[i])
			lines[j] = strings.TrimSpace(lines[j])
		}

		a := lines[i]
		b := lines[j]

		// Флаг выборки колоннки
		if fColumn > 0 {
			fieldsA := strings.Fields(a)
			fieldsB := strings.Fields(b)

			if fColumn > len(fieldsA) || fColumn > len(fieldsB) {
				return false
			}

			a = fieldsA[fColumn-1]
			b = fieldsB[fColumn-1]
		}

		// Флаг сортировка в обратном проядке
		if fReverse {
			a, b = b, a
		}

		// Флаг сортировка по числовому значению
		if fNumeric {
			numA, _ := strconv.Atoi(a)
			numB, _ := strconv.Atoi(b)

			// Флаг проверка отсортированы ли данные
			if fSort {
				if numA < numB {
					return true
				} else {
					isSorted = false
					return false
				}
			}
			return numA < numB
		}

		// Флаг сортировки по месяцам
		if fMonths {
			// Флаг проверки отсортированы ли данные
			if fSort {
				if month[a] < month[b] {
					return true
				} else {
					isSorted = false
					return false
				}
			}
			return month[a] < month[b]
		}

		// Флаг проверка отсортированы ли данные
		if fSort {
			if a < b {
				isSorted = false
				return true
			} else {
				return false
			}
		}
		return a < b
	})

	// Флаг для того, чтобы не выводить повторяющиеся строки
	if fUnique {
		lines = removeDuplicates(lines)
	}

	return lines
}
