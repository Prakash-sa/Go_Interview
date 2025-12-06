package channels

import "sync"

// CounterWithMutex increments a shared counter using a mutex-protected critical section.
func CounterWithMutex(n int) int {
	var (
		mu    sync.Mutex
		count int
		wg    sync.WaitGroup
	)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	return count
}

// CounterWithChannel models a single owner goroutine that owns the state.
// Callers send increments and receive the updated value on the ack channel.
func CounterWithChannel(n int) int {
	type incReq struct {
		delta int
		ack   chan int
	}

	inc := make(chan incReq)
	go func() {
		val := 0
		for req := range inc {
			val += req.delta
			req.ack <- val
		}
	}()

	var last int
	for i := 0; i < n; i++ {
		ack := make(chan int, 1)
		inc <- incReq{delta: 1, ack: ack}
		last = <-ack
	}
	close(inc)
	return last
}
