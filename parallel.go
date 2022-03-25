package go_parallel

import (
	"fmt"
	"sync"
)

type ParallelCallback[T any] func(chan T, chan Result, id int, wg *sync.WaitGroup)

type Result[T any] struct {
	id  int
	val T
}

func Run_parallel [T any](n_workers int, in chan T, out chan Result, Worker ParallelCallback[T]) {

	go func() {
		wg := sync.WaitGroup{}
		defer close(out) // close the output channel when all tasks are completed
		for id := 0; id < n_workers; id++ {
			wg.Add(1)
			go Worker(in, out, id, &wg)
		}
		wg.Wait() // wait for all workers to complete their tasks *and* trigger the -differed- close(out)
	}()
}

/*
// Example of use:
const (
	NW = 8
)

func Worker(in chan int, out chan Result, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range in {
		item *= 2 // returns the double of the input value (Bogus handling of data)
		out <- Result{id, item}
	}
}

func main() {

	in := make(chan int)
	out := make(chan Result)

	go func() {
		defer close(in)
		for i := 0; i < 10; i++ {
			in <- i
		}
	}()

	Run_parallel(NW, in, out, Worker)

	for item := range out {
		fmt.Printf("From out [%d]: %d\n", item.id, item.val)
	}

	println("- - - All done - - -")
}
*/
