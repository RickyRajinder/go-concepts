package main

import (
	"fmt"
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
	done := make(chan struct{})
	defer close(done)
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		c := produce(done, arr[i])
		out1 := consume(done, c)
		out2 := consume(done, c)
		out := merge(done, out2, out1)
		fmt.Println(<-out)
	}


}