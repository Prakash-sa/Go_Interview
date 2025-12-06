package timer

import "time"

// Rate limiting (Ticker / time.After)
// Notes: Ticker for sustained pacing; time.After for one-shot timeouts; always Stop() tickers to avoid leaks.

func RateLimiting(requests []int, handle func(int)) {
	tick := time.NewTicker(50 * time.Millisecond)
	defer tick.Stop()

	for _, r := range requests {
		<-tick.C // tick gates the pace
		handle(r)
	}
}
