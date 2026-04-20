package main

import "strings"

func article(word string) string {
	words := strings.Fields(word)
	for b := 0; b < len(words); b++ {
		vowels := "AEIOUHaeiouh"

		if words[b] == "A" && strings.ContainsAny(vowels, string(words[b+1][0])) {
			words[b] = "An"
		}

		if words[b] == "a" && strings.ContainsAny(vowels, string(words[b+1][0])) {
			words[b] = "an"
		}

		if words[b] == "An" && !strings.ContainsAny(vowels, string(words[b+1][0])) {
			words[b] = "A"
		}

		if words[b] == "an" && !strings.ContainsAny(vowels, string(words[b+1][0])) {
			words[b] = "a"
		}
	}
	return strings.Join(words, " ")
}
