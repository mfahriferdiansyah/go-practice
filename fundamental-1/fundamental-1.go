package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Waiting for input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput := scanner.Text()

	var answer = make(map[string]int)
	for _, char := range stringInput {
		fmt.Println(string(char))

		_, ok := answer[string(char)]
		if !ok {
			answer[string(char)] = 1
		} else {
			answer[string(char)] += 1
		}
	}
	fmt.Println(answer)
}
