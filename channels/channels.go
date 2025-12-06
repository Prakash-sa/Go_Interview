package channels

import (
	"strconv"
	"time"
)

// BasicSendReceive shows the synchronization of an unbuffered channel.
func BasicSendReceive() string {
	ch := make(chan string)
	go func() {
		ch <- "data"
	}()
	return <-ch
}

// FireAndForget spins a goroutine without waiting for it.
func FireAndForget(fn func()) {
	go fn()
}

// UnbufferedAndBuffered demonstrates blocking semantics for both channel types.
func UnbufferedAndBuffered() (unbuffered, buffered int) {
	ch := make(chan int)
	buf := make(chan int, 2)

	go func() {
		ch <- 42 // blocks until receiver is ready
		buf <- 7 // buffered send if capacity available
		buf <- 9 // still room because cap=2
		close(buf)
	}()

	unbuffered = <-ch
	buffered = <-buf
	return unbuffered, buffered
}

// Multiplexing uses select to handle whichever channel is ready first.
func Multiplexing() string {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- 2
	}()

	select {
	case v := <-ch1:
		return "ch1:" + strconv.Itoa(v)
	case w := <-ch2:
		return "ch2:" + strconv.Itoa(w)
	case <-time.After(200 * time.Millisecond):
		return "timeout"
	}
}

// ClosingSignal broadcasts completion by closing a channel.
func ClosingSignal() bool {
	done := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		return true
	case <-time.After(200 * time.Millisecond):
		return false
	}
}

// NonBlockingSend attempts a best-effort send using select default.
func NonBlockingSend(ch chan<- int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}
