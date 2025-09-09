package pipeline

import "context"

// Pipeline pattern (stages + cancellation)
// Notes: Each stage returns a read-only channel;
// caller ranges to drain; use ctx to avoid goroutine leaks when early-exiting.

func gen(ctx context.Context, nums ...int) <-chan int {
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

func sq(ctx context.Context, in <-chan int) <-chan int {
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

func runPipeline(ctx context.Context) []int {
    c := gen(ctx, 1, 2, 3, 4)
    out := sq(ctx, c)
    var res []int
    for v := range out {
        res = append(res, v)
    }
    return res
}
