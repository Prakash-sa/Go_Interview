# Go Interview

<div align="center">
	<img src="https://kislayverma.com/wp-content/uploads/2020/07/gopher-go.jpg" alt="Go Logo" width="350" />
	<br>
	<b>
		<span style="font-size:1.5em;">#goroutines &nbsp; #channels &nbsp; #concurrency &nbsp; #pipeline &nbsp; #context &nbsp; #timer</span>
	</b>
</div>

## Features

- **Garbage Collector** : GC is completely seamless in Go. There is no way for developers to tinker with it – no finalize() etc, it just works.
- **In built server packages** : Golang is wise to today’s API heavy world of today and give inbuilt package to start different types of servers. Firing up an HTTP server is about 5 lines of code.
- **Ground Up Concurrency** : Concurrency management is baked right into the language, and the importance of this is impossible to exaggerate. Developers can build scalable concurrent systems without extra frameworks or programming paradigm shifts (as you would have to do in most languages when going from sync to async).
- **Circular Dependency detection** : This was just awesome. Anyone who has dealt with keeping dependency graphs straight in large organizations will appreciate this. The compiler actually FAILS the build if it detects circular dependencies. This forces developers to be very careful about how the code is structured and what is upstream and what is downstream. This might cost some extra time during design, but definitely leads to more maintainable software in the long run.

## Important Notes

- Go’s concurrency model is based on goroutines and channels—master these for idiomatic Go.
- Always use context for cancellation, timeouts, and deadlines in concurrent code.
- Avoid sharing mutable state; prefer communication over channels.
- Use WaitGroups, errgroup, and select for robust coordination and error handling.
- Read the official Go blog and Rob Pike’s talks for deep insights.

- In programming, concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations. Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.

### Go Routines

![Go Routines](./routines.png)

## Common pitfalls & best practices

- **Goroutine leaks** : always have a way to stop goroutines—propagate ctx, close input channels, or use done signals. Pipelines must stop downstream stages on cancellation.

- **Range-and-close** : a receiver can range a channel until the sender closes it; don’t send after closing.

- **Backpressure** : pick buffer sizes intentionally; use semaphores to bound concurrency.

- **Timeouts** : prefer context.WithTimeout over sprinkled time.After in every select, to centralize control and avoid leaks.

- **Select fairness** : selection among ready cases is pseudo-random; don’t rely on case ordering for fairness.

- **Keep interfaces small**; compose rather than create “god” interfaces.

- **Don’t fight the model** : communicate by channels; share memory via synchronization when that’s simpler. Rob Pike’s talks & Go blog are gold.

# Links

- [Concurrency Patters](https://github.com/lotusirous/go-concurrency-patterns)
- [Go Interview](https://github.com/Devinterview-io/golang-interview-questions)
