package main

import (
	"regexp"
	"strconv"
	"strings"
)

func ToUpperLastWord(text []string) []string {

	if len(text) < 2 {
		return text
	}

	b := len(text) - 1

	if text[b] == "(up)" {

		text[b-1] = strings.ToUpper(text[b-1])
		text = text[:b]
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

func fixQuotes(text string) string {
	p1 := regexp.MustCompile(`\s+'`)
	p2 := regexp.MustCompile(`'\s`)

	text = p1.ReplaceAllString(text, "'")
	text = p2.ReplaceAllString(text, "'")

	text = strings.ReplaceAll(text, ":", ": ")

	return text
}

func fixPunctuation(text string) string {

	p1 := regexp.MustCompile(`\s+([?.,:;!]+)`)
	p2 := regexp.MustCompile(`([?.,:;!]+)\s+`)

	text = p1.ReplaceAllString(text, "$1 ")
	text = p2.ReplaceAllString(text, "$1")
	result := strings.Fields(text)
	return strings.Join(result, " ")
}

func fixarticles(s string) string {
	words := strings.Fields(s)
	vowels := "aeiouhAEIOUH"

	for i := 0; i < len(words); i++ {
		if i < len(words) {
			word := words[i]
			if word == "a" && strings.ContainsAny(vowels, string(string(words[i+1][0]))) {
				words[i] = "an"
			} else if word == "A" && strings.ContainsAny(vowels, string(string(words[i+1][0]))) {
				words[i] = "An"
			} else if word == "an" && !strings.ContainsAny(vowels, string(string(words[i+1][0]))) {
				words[i] = "a"
			}
		}

	}

	return strings.Join(words, " ")

}
