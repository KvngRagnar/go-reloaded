package main

import (
	"regexp"
)

func punctQuote(text string) string {

	p1 := regexp.MustCompile(`\s*([?.,;:!]+)`)
	text = p1.ReplaceAllString(text, "$1")

	p2 := regexp.MustCompile(`([?.,;:!]+)\s*`)
	text = p2.ReplaceAllString(text, "$1 ")

	p3 := regexp.MustCompile(`'\s*(.*?)\s*'`)
	text = p3.ReplaceAllString(text, `'$1' `)

	p4 := regexp.MustCompile(`"\s*(.*?)\s*"`)
	text = p4.ReplaceAllString(text, `"$1" `)

	return text
}
