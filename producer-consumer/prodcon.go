package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

//producer channel to pass input to consumer
func produce(done <-chan struct{}, nums ...int) <-chan int{
	out := make(chan int, len(nums))
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
				case out <- n:
				case <- done:
					return
			}
		}
	}()
	return out
}

//consumer channel to square value from producer
func consume(done <-chan struct{}, in <- chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
				case out <- n * n:
				case <- done:
					return
			}
		}
	}()
	return out
}

//merge two or more goroutines
func merge(done <-chan struct{}, cs ...<- chan int) <- chan int {
	var waitgroup sync.WaitGroup
	out := make(chan int)

	output := func(c <- chan int){
		defer waitgroup.Done()
		for n:= range c {
			select {
				case out <- n:
				case <- done:
					return
			}
		}
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
	fptr := flag.String("fpath", "nums.txt", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	str := string(data)
	nums := strings.Split(str, "\n")

	done := make(chan struct{})
	defer close(done)

	for i := 0; i < len(nums); i++ {
		in, _ := strconv.Atoi(nums[i])
		c := produce(done, in)
		out1 := consume(done, c)
		out2 := consume(done, c)
		out := merge(done, out2, out1)
		fmt.Println(<-out)
	}


}