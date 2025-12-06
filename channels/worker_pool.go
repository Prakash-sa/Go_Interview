package channels

import (
	"context"
	"sync"
)

// Fan-in (merge many into one)
func fanIn[T any](chs ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(chs))
	for _, ch := range chs {
		ch := ch
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Fan-out / worker pool
// Pattern: N workers reading from one jobs channel, sending results

// Notes: This is the canonical worker-pool (fan-out) with context for cancellation and a single results stream (fan-in).
// Ensure you consume all results or close result channel when done.

// Notes: One producer → N workers (fan-out) → one results stream (fan-in).
// Ensure you fully drain output or close appropriately to avoid leaks.

type Job struct{ ID int }
type Result struct {
	ID  int
	Err error
}

func worker(ctx context.Context, jobs <-chan Job, out chan<- Result) {
	for {
		select {
		case <-ctx.Done():
			return
		case j, ok := <-jobs:
			if !ok {
				return
			}
			// process j
			out <- Result{ID: j.ID, Err: nil}
		}
	}
}

func runPool(ctx context.Context, n int, inputs []Job) []Result {
	if n <= 0 {
		n = 1
	}

	jobs := make(chan Job)
	out := make(chan Result)

	// start workers
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			worker(ctx, jobs, out)
		}()
	}

	// feed jobs
	go func() {
		defer close(jobs)
		for _, j := range inputs {
			select {
			case <-ctx.Done():
				return
			case jobs <- j:
			}
		}
	}()

	// close out after all workers exit
	go func() {
		wg.Wait()
		close(out)
	}()

	// collect results
	results := make([]Result, 0, len(inputs))
	for r := range out {
		results = append(results, r)
		if len(results) == len(inputs) {
			return results
		}
	}
	return results
}
