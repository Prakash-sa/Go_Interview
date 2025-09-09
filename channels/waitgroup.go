package channels

// sync.WaitGroup for lifecycle coordination

// Notes: Only call Add before goroutines start
// pair each goroutine with a defer wg.Done(). Use in tandem with context for cancellation.

func waitgroups() {
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, t := range tasks {
		t := t
		go func() {
			defer wg.Done()
			_ = work(t)
		}()
	}
	wg.Wait()
}
