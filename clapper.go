package clapper

import (
	"os"
	"strings"
)

// Clappify a string
func Clap(victim string) (result string) {
	// Get the length of the clap
	length := len(strings.Split(victim, " "))

	if length == 1 {
		// When we have one word clap between letters
		result = "👏 "
		for _, char := range os.Args[1] {
			result = result + string(char) + " 👏 "
		}
	} else {
		// If we have a sentence then clap between words
		str := strings.Join(os.Args[1:], " ")
		result = strings.Replace(str, " ", " 👏 ", -1)
		result = "👏 " + result + " 👏"
	}

	// Trim any nasty whitespace and display result
	result = strings.TrimSpace(result)
	return
}

