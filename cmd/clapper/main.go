package main

import (
	"strings"
	"os"
	"fmt"
	"github.com/rcole5/clapper"
)

func main() {
	if len(os.Args) < 2 {
		// Check if we have anything to clap on
		fmt.Println("ERR: Not enough args.")
		os.Exit(2)
	}

	// Execute the command
	victim := strings.Join(os.Args[1:], " ")
	result := clapper.Clap(victim)

	// Display the result
	fmt.Println(result)
}
