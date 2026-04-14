package main

import (
	"regexp"
	"strconv"
	"strings"
)

func cases(text string) string {
	words := strings.Fields(text)
	for b := len(words) - 1; b >= 0; b-- {

		if words[b] == "(low)" && b > 0 {
			words[b-1] = strings.ToLower(words[b-1])
			words = append(words[:b], words[b+1])
		}

		if words[b] == "(up)" && b > 0 {
			words[b-1] = strings.ToUpper(words[b-1])
			words = append(words[:b], words[b+1])
		}

		if words[b] == "(cap)" && b > 0 {
			words[b-1] = strings.ToUpper(string(words[b-1][0]) + strings.ToLower(words[b-1][1:]))
			words = append(words[:b], words[b+1])
		}

		if strings.HasPrefix(words[b], "(up,") && b+1 < len(words) {
			index := strings.TrimSuffix(words[b+1], ")")
			num, err := strconv.Atoi(index)
			if err != nil {
				continue
			}
			for j := 1; j <= num; j++ {
				if b-j >= 0 {
					words[b-j] = strings.ToUpper(words[b-j])
				}
			}
			words = append(words[:b], words[b+2:]...)
		}

		if strings.HasPrefix(words[b], "(low,") && b+1 < len(words) {
			index := strings.TrimSuffix(words[b+1], ")")
			num, err := strconv.Atoi(index)
			if err != nil {
				continue
			}
			for j := 1; j <= num; j++ {
				if b-j >= 0 {
					words[b-j] = strings.ToLower(words[b-j])
				}
			}
			words = append(words[:b], words[b+2])
		}

		if strings.HasPrefix(words[b], "(cap,") && b+1 < len(words) {
			index := strings.TrimSuffix(words[b+1], ")")

			num, err := strconv.Atoi(index)
			if err != nil {
				continue
			}

			for j := 1; j <= num; j++ {
				if b-j >= 0 {
					words[b-j] = strings.ToUpper(string(words[b-j][0])) + strings.ToLower(words[b-j][1:])
				}
			}

			words = append(words[:b], words[b+2:]...)
		}

	}
	return strings.Join(words, " ")
}


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
	// p1 := regexp.MustCompile(`\s+'`)
	// p2 := regexp.MustCompile(`'\s`)

	// text = p1.ReplaceAllString(text, "'")
	// text = p2.ReplaceAllString(text, "'")

	// text = strings.ReplaceAll(text, ":", ": ")

	// return text

	words := strings.Split(text, "'")
	for i := range words {
		if i%2 == 1 {
			words[i] = strings.TrimSpace(words[1])
		}
	}
	return strings.Join(words, "'")
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
