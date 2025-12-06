package channels

import "sync"

// ProcessWithSemaphore bounds concurrency using a buffered channel as a semaphore.
func ProcessWithSemaphore[T any](items []T, limit int, work func(T) T) []T {
	if limit <= 0 {
		limit = 1
	}

	sem := make(chan struct{}, limit)
	var wg sync.WaitGroup
	out := make([]T, len(items))

	for i, item := range items {
		wg.Add(1)
		sem <- struct{}{} // acquire
		go func(idx int, val T) {
			defer wg.Done()
			defer func() { <-sem }() // release
			out[idx] = work(val)
		}(i, item)
	}

	wg.Wait()
	return out
}
