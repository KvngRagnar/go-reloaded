package main

import (
	"strconv"
	"strings"
)

func cases(text string) string {

	words := strings.Fields(text)
	for b := len(words) - 1; b >= 0; b-- {

		if words[b] == "(bin)" && b > 0 {

			num, err := strconv.ParseInt(words[b-1], 2, 64)
			if err == nil {
				words[b-1] = strconv.FormatInt(num, 10)
			}

			words = append(words[:b], words[b+1:]...)
		}

		if words[b] == "(hex)" && b > 0 {

			num, err := strconv.ParseInt(words[b-1], 16, 64)
			if err == nil {
				words[b-1] = strconv.FormatInt(num, 10)
			}

			words = append(words[:b], words[b+1:]...)
		}

		if words[b] == "(up)" && b > 0 {
			words[b-1] = strings.ToUpper(words[b-1])
			words = append(words[:b], words[b+1:]...)
		}

		if words[b] == "(low)" && b > 0 {
			words[b-1] = strings.ToLower(words[b-1])
			words = append(words[:b], words[b+1:]...)
		}

		if words[b] == "(cap)" && b > 0 {
			words[b-1] = strings.ToUpper((string)(words[b-1][0])) + strings.ToLower(words[b-1][1:])
			words = append(words[:b], words[b+1:]...)
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
			words = append(words[:b], words[b+2:]...)

		}

		if strings.HasPrefix(words[b], "(cap,") && b+1 < len(words) {
			index := strings.TrimSuffix(words[b+1], ")")
			num, err := strconv.Atoi(index)
			if err != nil {
				continue
			}
			for j := 1; j <= num; j++ {
				if b-j >= 0 {
					words[b-j] = strings.ToUpper((string)(words[b-j][0])) + strings.ToLower(words[b-j][1:])
				}
			}
			words = append(words[:b], words[b+2:]...)
		}

	}
	return strings.Join(words, " ")
}
