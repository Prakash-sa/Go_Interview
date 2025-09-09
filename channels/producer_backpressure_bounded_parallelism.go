package channels

// Notes: A buffered channel as a counting semaphore gives explicit concurrency caps 
// (useful for API rate limits or CPU-bound ops).
func main(){
	// semaphore with capacity K limits concurrency
	sem := make(chan struct{}, 8)

	for _, item := range items {
		item := item
		sem <- struct{}{}        // acquire
		go func() {
			defer func() { <-sem }() // release
			_ = process(item)
		}()
	}

	// wait for all permits to drain
	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}
}


