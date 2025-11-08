package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	for true {
		fmt.Print("Pokedex >")

		scanner.Scan()
		input = scanner.Text()
		cl_input := cleanInput(input)
		fmt.Println("Your command was:", cl_input[0])
	}
}
