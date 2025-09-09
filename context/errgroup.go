package context

import (
	"context"

	"golang.org/x/sync/errgroup"
)


// errgroup to run tasks and collect the first error
// Notes: errgroup cancels sibling goroutines on first error through the derived ctx. Great for parallel I/O. 
// (Standard practice alongside pipelines.) 

func runAll(ctx context.Context, urls []string) error {
    g, ctx := errgroup.WithContext(ctx)
    for _, u := range urls {
        u := u
        g.Go(func() error {
            // do work; return error to cancel siblings
            if err := fetch(ctx, u); err != nil {
                return err
            }
            return nil
        })
    }
    return g.Wait() // first error or nil
}

