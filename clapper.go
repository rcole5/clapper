package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var result string
	if len(os.Args) < 2 {
		// Check if we have anything to clap on
		fmt.Println("ERR: Not enough args.")
		os.Exit(2)
	} else if len(os.Args) == 2 {
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
	fmt.Println(result)
}
