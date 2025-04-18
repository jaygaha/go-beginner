# Go Scheduler

This guide introduces the Go scheduler, a key component of the Go runtime that manages goroutines, enabling efficient concurrency. We'll cover the basics, how it works, and provide step-by-step tutorials to help beginners understand and use the scheduler effectively.

## What is the Go Scheduler?

The Go scheduler is part of the Go runtime and is responsible for managing goroutines, which are lightweight threads. Unlike OS threads, goroutines are multiplexed onto a smaller number of OS threads by the scheduler, making concurrency efficient and scalable. The scheduler ensures that goroutines are executed, paused, or resumed as needed, balancing CPU usage and minimizing idle time.

### Key Concepts

- `Goroutine`: A lightweight thread managed by the Go runtime.
- `GOMAXPROCS`: The number of OS threads available to the scheduler (defaults to the number of CPU cores).
- `M:N Scheduling`: Maps many goroutines (M) to a few OS threads (N).
- `Run Queue`: Each OS thread has a queue of goroutines waiting to run.
- `Preemption`: The scheduler can interrupt long-running goroutines to ensure fairness.
- `Work Stealing`: Idle threads can "steal" goroutines from busy threads to balance workload.

### How Does the Scheduler Work?

- `Goroutine Creation`: When you use the go keyword, a new goroutine is created and added to a run queue.
- `Thread Allocation`: The scheduler assigns goroutines to available OS threads (controlled by GOMAXPROCS).
- `Execution`: The thread executes the goroutine until it completes, blocks (e.g., I/O), or is preempted.
- `Context Switching`: If a goroutine blocks (e.g., waiting for I/O), the scheduler swaps it out and runs another goroutine on the same thread.
- `Work Stealing`: If a thread runs out of goroutines, it steals from another thread’s run queue.

### Tips for Working with the Go Scheduler

- `Avoid Blocking Main`: Use synchronization mechanisms like channels or `sync.WaitGroup` to prevent main from exiting prematurely.
- `Tune GOMAXPROCS`: Experiment with `runtime.GOMAXPROCS` for CPU-bound tasks, but the default is usually fine.
- `Minimize Blocking`: Use non-blocking operations or channels to keep goroutines responsive.
- `Monitor Goroutines`: Use `runtime.NumGoroutine()` to debug and check how many goroutines are active.
- `Leverage Preemption`: Ensure loops yield periodically (e.g., with `runtime.Gosched()`) in long-running tasks to avoid starving other goroutines.

### Common Pitfalls

- `Race Conditions`: Goroutines accessing shared data may cause races. Use sync.Mutex or channels for safe access.
- `Premature Exit`: Without proper synchronization, main may exit before goroutines finish.
- `Overusing Goroutines`: Creating too many goroutines can strain the scheduler. Use worker pools for heavy workloads.

### Examples

1. **Basic Goroutine Creation** [Check](21_1_basic/main.go)
2. **GOMAXPROCS** [Check](21_2_gomaxprocs/main.go)
3. **Blocking Operations** [Check](21_3_blocking_operations/main.go)
4. **Channels** [Check](21_4_channels/main.go)

## Further Reading

- Read the [Go Concurrency Patterns](https://go.dev/talks/2012/concurrency.slide) by Rob Pike.
- Experiment with `sync.WaitGroup` for more complex synchronization.
- Explore the `runtime` package for advanced scheduler control.
- Use tools like `go tool trace` to visualize scheduler behavior.
