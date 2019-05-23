package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("hello, world\n")
	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter Text: ")
		text, _ := scanner.ReadString('\n')
		if len(text) > 0 {
			fmt.Println(text)
		}
	}
}
