package context

import (
	"context"
	"time"
)

// cancellation & timeout with context
// Prefer context over ad-hoc “done” channels for deadlines and propagation; always defer cancel() in the creator.

func main(){
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	select {
	case <-doWork(ctx):   // returns a channel closed on completion
	case <-ctx.Done():    // deadline exceeded or canceled
		// handle timeout
	}
}