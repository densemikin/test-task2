package main

import (
	"fmt"
	"strings"
)

const vowels = "aeiou"
const prefix = "ay"
const sep = " "

func translate(in string) string {

	var indexOfFirstVowel = 0

	for _, charValue := range strings.Split(in, "") {
		if strings.Contains(vowels, charValue) || strings.Contains(vowels, strings.ToLower(charValue)) {
			if indexOfFirstVowel == 0 {
				return in + prefix
			} else {
				break
			}
		} else {
			indexOfFirstVowel++
		}
	}
	return in[indexOfFirstVowel:] + in[:indexOfFirstVowel] + prefix
}

func translateSentences(inputSentence string) string {

	words := strings.Split(inputSentence, " ")
	var wordsList []string

	for _, word := range words {
		wordsList = append(wordsList, translate(word))
	}
	return strings.Join(wordsList, sep)
}

func translateToPigLatin() {
	fmt.Println(translateSentences("eat eat") == "eatay eatay")
	fmt.Println(translateSentences("EAT EAT") == "EATay EATay")
	fmt.Println(translateSentences("omelet") == "omeletay")
	fmt.Println(translateSentences("explain") == "explainay")
	fmt.Println(translateSentences("pig") == "igpay")
	fmt.Println(translateSentences("latin") == "atinlay")
	fmt.Println(translateSentences("will") == "illway")
	fmt.Println(translateSentences("butler") == "utlerbay")
	fmt.Println(translateSentences("duck") == "uckday")
	fmt.Println(translateSentences("me") == "emay")
}