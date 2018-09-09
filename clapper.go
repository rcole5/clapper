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
		result = "👏 "
		for _, char := range arr {
			result = result + string(char) + " 👏 "
		}
	} else {
		// If we have a sentence then clap between words
		str := strings.Join(arr, " ")
		result = strings.Replace(str, " ", " 👏 ", -1)
		result = "👏 " + result + " 👏"
	}

	// Trim any nasty whitespace and display result
	result = strings.TrimSpace(result)
	return
}

