package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	text := string(data)
	result := textProcessor(text)

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func textProcessor(text string) string {
	text = cases(text)
	text = punctQuote(text)
	text = article(text)

	return text + "\n"
}
