package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Mode()&os.ModeNamedPipe != 0 {
		r := bufio.NewReader(os.Stdin)
		var input []byte
		for {
			in, err := r.ReadByte()
			if err != nil && err == io.EOF {
				break
			}
			input = append(input, in)
		}
		for i := 0; i < len(input); i++ {
			fmt.Printf("%c", input[i])
		}
	}
}
