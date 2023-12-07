package main

import (
	"reflect"
	"testing"
)

func TestAnagram(t *testing.T) {
	words := []string{"Пятак", "пЯтка", "анаграммынету", "ТЯПКА"}
	expected := &map[string]*[]string{
		"пятак": {"пятак", "пятка", "тяпка"},
	}
	result := anagram(&words)
	if !reflect.DeepEqual(result, expected) {
		t.Error("Результат неправильный")
	}

}
