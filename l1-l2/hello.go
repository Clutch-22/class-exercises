// Every executable Go Program should contain a package called main.
// This tells the Go compiler to compile the package into an executable
// program rather than a shared library.
package main

import (
	"fmt"
	"os"
)

// The entry point of a Go program should be the main function of main package.
// When the executable is run, main() is automatically called.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func searchForWord(filepath string, target string) {

	dat, err := os.ReadFile(string(filepath))
	check(err)
	fmt.Print(string(dat))
	input := string(filepath)
	foundString := string(target)
	iterator := 1

	for iterator <= len(input)-1 {
		if foundString == input[iterator-1:iterator+1] {
			fmt.Printf("found ", foundString, " @ %v\n", iterator-1)
		}
		iterator += 1
	}
}

func main() {

	fmt.Println("An oasis of love and friendship.")
	searchForWord("dictionary.txt", "re")

}
