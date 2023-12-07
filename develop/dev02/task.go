package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Раскаповка токена
func unpackToken(token string) string {
	if len(token) == 1 {
		return token
	} else {
		symb := token[0:1]
		count, err := strconv.Atoi(token[1:])
		if err != nil {
			panic(err)
		}
		return strings.Repeat(symb, count)
	}
}

// Функция распаковки
func unpack(str string) (string, error) {
	if len(str) == 0 {
		return str, nil
	}

	// Создание регулярного выражения для разбивки на токены
	splitterRegexp := regexp.MustCompile(`\D{1}\d{0,}`)
	// Разбивка на строки
	tokens := splitterRegexp.FindAllString(str, -1)

	// После разбивки некоторые смволы могут не образовать токены
	// Это значит что исходная строка не может быть распакована =>

	if len(tokens) > 0 && strings.Join(tokens, "") == str {
		result := make([]string, len(tokens))
		for i, token := range tokens {
			result[i] = unpackToken(token)
		}
		return strings.Join(result, ""), nil
	} else {
		return "", errors.New("Ошибка")
	}
}

func main() {
	packed := "a3b3c4"
	unpacked, err := unpack(packed)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(unpacked)
}
