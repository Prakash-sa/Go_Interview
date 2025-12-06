package pipeline

// Simple two-stage pipeline: source -> square -> sink.
// Each stage owns its output channel and closes it when done to avoid leaks.

// SliceToChannel fan-outs a slice into a channel.
func SliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// Square squares incoming numbers and forwards results.
func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// Run composes the pipeline and collects results for convenience in examples/tests.
func Run(nums []int) []int {
	src := SliceToChannel(nums)
	sq := Square(src)

	var res []int
	for v := range sq {
		res = append(res, v)
	}
	return res
}
