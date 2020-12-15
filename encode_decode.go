package main

import (
	"fmt"
	"strings"
)

var vowelsToNumbers = map[string]string{
	"a": "1",
	"e": "2",
	"i": "3",
	"o": "4",
	"u": "5",
	"A": "1",
	"E": "2",
	"I": "3",
	"O": "4",
	"U": "5",
}

var numbersToVowels = map[string]string{
	"1": "a",
	"2": "e",
	"3": "i",
	"4": "o",
	"5": "u",
}


func replaceCharacters(in string, replaceTable map[string]string) string {
	out := in
	for fromChar, toChar := range replaceTable {
		out = strings.ReplaceAll(out, fromChar, toChar)
	}
	return out
}

func enncodeDecodeTestTask() {
	fmt.Println(replaceCharacters("QWEQWE", vowelsToNumbers) == "QW2QW2")
	fmt.Println(replaceCharacters("QW22QW22", numbersToVowels) == "QWeeQWee")
	s3 := "aaaqqqeessccvvbbrruuuioo"
	fmt.Println(replaceCharacters(replaceCharacters(s3, vowelsToNumbers), numbersToVowels) == s3)

}