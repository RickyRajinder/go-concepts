package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

//producer channel to pass input to consumer
func produce(nums ...int) <-chan int{
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

//consumer channel to square value from producer
func consume(in <- chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

//merge two or more goroutines
func merge(cs ...<- chan int) <- chan int {
	var waitgroup sync.WaitGroup
	out := make(chan int)

	output := func(c <- chan int){
		for n:= range c {
			out <- n
		}
		waitgroup.Done()
	}
	waitgroup.Add(len(cs))
	for _, c:= range cs {
		go output(c)
	}

	go func() {
		waitgroup.Wait()
		close(out)
	}()
	return out
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a Number: ")
		scanner.Scan()
		int, _ := strconv.Atoi(scanner.Text())
		c := produce(int)
		out1 := consume(c)
		out2 := consume(c)

		out := merge(out2, out1)
		fmt.Println(<-out)
	}
}