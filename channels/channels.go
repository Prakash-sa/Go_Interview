package channels

import (
	"fmt"
	"time"
)

func main(){
	myChannels:=make(chan string)
	go func(){
		myChannels<-"data"
	}()

	msg:=<-myChannels
	fmt.Println(msg)
}

// Fire-and-forgot goroutine
func fire_and_forgot_gorountine(){ 
	go func() {
		// do work concurrently
	}()
}


// Unbuffered and buffered channel
// Notes: Unbuffered channels synchronize sender/receiver
// buffered channels add queueing capacity (use to express backpressure). 
func channels(){
	// unbuffered: send blocks until the receiver is ready(synchronizes)
	ch:=make(chan int)

	// buffered: send blocks only when the buffer is full (backpressure)(asynchonous)
	buf:=make(chan int,8)

	go func(){
		ch <- 42 // blocks until main receives
    	buf <- 7 // may not block if capacity available
	}()

	x := <-ch
	y := <-buf
	_ = x
	_ = y
}


// select for multiplexing
// select picks a ready case at random if multiple are ready; 
// use time.After for per-operation timeouts.
func multiplexing(){
	ch1:=make(chan int)
	ch2:=make(chan int)
	select{
	case v:=<-ch1:
		_=v // handle channel 1
	case w:=<-ch2:
		_=w // handle channel 2
	case <-time.After(200*time.Millisecond):
		// handle timeout fallback
	}
}

// Closing channels & signaling
// Broadcast “done” by closing a channel
done := make(chan struct{})

go func() {
    // ... when finished:
    close(done) // all receivers unblock
}()

select {
case <-done:
    // observed close
}


// Non-blocking sends/receives with select default
// Notes: Use sparingly; default makes the operation non-blocking.
// Good for loss-tolerant metrics or best-effort signals. 

select {
case ch <- v:
    // sent
default:
    // channel full: drop, log, or apply backpressure
}


