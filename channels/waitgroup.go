package channels

import "sync"

// WaitGroups coordinates goroutines and waits for them to finish.
func WaitGroups(tasks []func()) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, t := range tasks {
		task := t
		go func() {
			defer wg.Done()
			task()
		}()
	}
	wg.Wait()
}
