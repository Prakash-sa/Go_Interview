package timer

// Timers vs tickers vs time.AfterFunc
// Notes: Timer is single-shot; Ticker is periodic.
// Drain channels on stop when necessary to avoid stray wakeups.

import (
	"context"
	"time"
)

func OneShotTimer(ctx context.Context, d time.Duration) bool {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-timer.C: // fired once
		return true
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C // drain if already fired
		}
		return false
	}
}
