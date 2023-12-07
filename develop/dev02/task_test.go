package main

import "testing"

func TestUnpack(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abcd",
			expected: "abcd",
		},
		{
			input:    "45",
			expected: "",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "a10df5cvs1",
			expected: "aaaaaaaaaadfffffcvs",
		},
	}
	for _, testCase := range testCases {
		result, _ := unpack(testCase.input)
		if result != testCase.expected {
			t.Errorf("Неверный результат. Ожидалось: %s, получено %s", testCase.expected, result)
		}
	}
}
