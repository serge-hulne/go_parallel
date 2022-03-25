# go_parallel

Go library (wit generics) to run worker process in parallel (concurrently)

```
// Example of use

package main

import (
	"fmt"
	"sync"

	. "github.com/serge-hulne/go_parallel"
)

// Example of use:
const (
	NW         = 8
	BufferSize = 1
)

func Worker(in chan int, out chan Result[int], id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range in {
		item *= 2 // returns the double of the input value (Bogus handling of data)
		out <- Result[int]{id, item}
	}
}

func main() {

	// in and out channels:
	in := make(chan int, BufferSize)
	out := make(chan Result[int])

	// Populate in channel (send data to input stream)
	go func() {
		defer close(in)
		for i := 0; i < 10; i++ {
			in <- i
		}
	}()

	// Run tasks (Worker func) in parallel:
	Run_parallel(NW, in, out, Worker)

	// Display results:
	for item := range out {
		fmt.Printf("From out [%d]: %d\n", item.Id, item.Value)
	}
	println("- - - All done - - -")
}

```

