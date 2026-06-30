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

func main() {

	dat, err := os.ReadFile("dictionary.txt")
	check(err)
	fmt.Print(string(dat))

	fmt.Println("An oasis of love and friendship.")

	//input := "Due to your wasting of company resources by hogging the genetic recombinator, you have been selected to participate in an experimental biological enhancement program. All your debt is cleared and from now on your body will regenerate by itself. Can't say I envy you, though."
	input := string(dat)

	foundString := "re"

	iterator := 1

	for iterator <= len(input)-1 {
		if foundString == input[iterator-1:iterator+1] {
			fmt.Printf("found re @ %v\n", iterator-1)
		}
		iterator += 1
	}
}
