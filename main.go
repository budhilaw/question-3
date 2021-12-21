package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := `((halo semua (123) (123))`

	fmt.Println("Question no. 3: ", findFirstStringInBracket(s))
	fmt.Println("Answer: ", simpleFunction(s))
}

func simpleFunction(str string) string {
	/**
	 * \( -> match character (
	 * (.*?) -> match any characters inside capturing group
	 * [\s\S\d\D\w\W] -> match any characters
	 * \) -> match character )
	 */
	regex := regexp.MustCompile(`\((.*?)[\s\S\d\D\w\W]\)`)
	res := regex.FindStringSubmatch(str)
	return res[1]
}

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
	return ""
}
