package pipeline

import "context"

// Pipeline pattern (stages + cancellation).
// Notes: Each stage returns a read-only channel; caller drains it.
// Always respect ctx cancellation to prevent goroutine leaks.

func Gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

func SquareWithContext(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			n2 := n * n
			select {
			case <-ctx.Done():
				return
			case out <- n2:
			}
		}
	}()
	return out
}

func RunWithContext(ctx context.Context, nums ...int) []int {
	c := Gen(ctx, nums...)
	out := SquareWithContext(ctx, c)

	var res []int
	for v := range out {
		res = append(res, v)
	}
	return res
}
