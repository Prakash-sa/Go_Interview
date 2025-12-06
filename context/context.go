package contextdemo

import (
	"context"
	"net/http"
	"time"
)

// WorkEmulator returns a channel closed when the fake work finishes or ctx cancels.
func WorkEmulator(ctx context.Context, d time.Duration) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		select {
		case <-time.After(d):
		case <-ctx.Done():
		}
	}()
	return done
}

// WithTimeout demonstrates cancellation handling using context.
func WithTimeout() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-WorkEmulator(ctx, 10*time.Millisecond):
		return true // completed before timeout
	case <-ctx.Done():
		return false // deadline exceeded
	}
}

// HTTPWithContext shows how HTTP requests inherit deadlines and cancellations.
func HTTPWithContext(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
