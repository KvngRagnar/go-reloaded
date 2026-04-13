package main

import (
	"strconv"
	"strings"
)

func ToUpperLastWord(text []string) []string {

	if len(text) < 2 {
		return text
	}

	lastIndex := len(text) - 1

	if text[lastIndex] == "(up)" {

		text[lastIndex-1] = strings.ToUpper(text[lastIndex-1])
		text = text[:lastIndex]
	}
	return text
}

func convertNumbers(word []string) []string {

	for i := 1; i < len(word); i++ {
		switch word[i] {

		case "(hex)":
			value, _ := strconv.ParseInt(word[i-1], 16, 64)
			word[i-1] = strconv.FormatInt(value, 10)
			word = append(word[:i], word[i+1:]...)

		case "(bin)":
			value, _ := strconv.ParseInt(word[i-1], 2, 64)
			word[i-1] = strconv.FormatInt(value, 10)
			word = append(word[:i], word[i+1:]...)
		}
	}
	return word
}
