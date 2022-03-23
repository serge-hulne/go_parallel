# go_parallel

Go library (wit generics) to run worker process in parallel (concurrently)

```
// Example of use

const (
	NW = 8
)

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
```

