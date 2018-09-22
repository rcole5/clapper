package clapper

import (
	"strings"
)

// Clappify a string
func Clap(victim string) (result string) {
	// Get the length of the clap
	arr := strings.Split(victim, " ")
	length := len(arr)

	if length == 1 {
		// When we have one word clap between letters
		return LetterClap(victim)
	} else {
		// If we have a sentence then clap between words
		return WordClap(victim)
	}
}

// Clap the letters
func LetterClap(victim string) (result string) {
	arr := strings.Split(victim, " ")
	result = "👏 "
	for _, char := range arr {
		result = result + string(char) + " 👏 "
	}
	return strings.TrimSpace(result)
}

// Clap the words
func WordClap(victim string) (result string) {
	arr := strings.Split(victim, " ")
	str := strings.Join(arr, " ")
	result = strings.Replace(str, " ", " 👏 ", -1)
	result = "👏 " + result + " 👏"
	return strings.TrimSpace(result)
}
