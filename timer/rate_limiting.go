package timer

import "time"

// Rate limiting (Ticker / time.After)
// Notes: Ticker for sustained pacing; time.After for one-shot timeouts; always Stop() tickers to avoid leaks.

func rate_limiting() {
	lim := time.NewTicker(50 * time.Millisecond)
	defer lim.Stop()

	for _, r := range requests {
		<-lim.C // tick gates the pace
		_ = handle(r)
	}
}
