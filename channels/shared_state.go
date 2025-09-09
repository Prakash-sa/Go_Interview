package channels


// Shared state: sync.Mutex vs channels
// Mutex: protect critical section

// Notes: Prefer channels to communicate, mutex to protect; design around ownership of data. 
// Choose the simpler tool; mutex is fine for small critical sections. 

var mu sync.Mutex
count := 0

inc := func() {
    mu.Lock()
    count++
    mu.Unlock()
}

// Channel-as-owner: single goroutine owns state
type incReq struct{ 
	delta int; 
	ack chan int 
}

func counter() (inc chan<- incReq, reads <-chan int) {
    inc = make(chan incReq)
    out := make(chan int)
    go func() {
        defer close(out)
        val := 0
        for req := range inc {
            val += req.delta
            req.ack <- val
        }
    }()
    return inc, out
}
