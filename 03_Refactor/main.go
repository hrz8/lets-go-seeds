package main

import (
	"fmt"
	"strings"
)

func main() {
	res := FindFirstStringInBracket("Nuruddin (Zayn Malik)")
	fmt.Println(res)
}

func FindFirstStringInBracket(str string) string {
	indexFirstBracketFound := strings.Index(str, "(")
	if len(str) == 0 || indexFirstBracketFound < 0 {
		return ""
	}
	strRunes := []rune(str)
	wordsAfterFirstBracket := string(strRunes[indexFirstBracketFound:len(str)])
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound < 0 {
		return ""
	}
	afterBracketRunes := []rune(wordsAfterFirstBracket)
	return strings.Split(string(afterBracketRunes[1:indexClosingBracketFound]), " ")[0]
}
