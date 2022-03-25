package github.com/serge-hulne/go_parallel

import (
	"sync"
)

type ParallelCallback[T any] func(chan T, chan Result[T], int, *sync.WaitGroup)

type Result[T any] struct {
	Id    int
	Value T
}

func Run_parallel[T any](n_workers int, in chan T, out chan Result[T], Worker ParallelCallback[T]) {
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
