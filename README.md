# go_parallel

Go library (wit generics) to run worker process in parallel (concurrently)

```
// Example of use

const (
	NW = 8
)

func Worker(in chan int, out chan Result[int], id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range in {
		item *= 2 // returns the double of the input value (Bogus handling of data)
		out <- Result[int]{id, item}
	}
}

func main() {
	in := make(chan int)
	out := make(chan Result[int])
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
```

