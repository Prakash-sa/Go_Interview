package timer

// Timers vs tickers vs time.AfterFunc
// Notes: Timer is single-shot; Ticker is periodic. 
// Drain channels on stop when necessary to avoid stray wakeups. 

func timer() {
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C: // fired once
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		} // drain if already fired
	}
}


