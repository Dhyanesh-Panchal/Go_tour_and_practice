# `runtime`

The runtime package in Go provides low-level access to the Go runtime system, including memory management, goroutines, scheduling, garbage collection, and system interactions.

## Go routine management and Scheduling

- ``runtime.GOMAXPROC(n int)`` GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting.
  - It defaults to the value of runtime.NumCPU. If n < 1, it does not change the current setting. 
[Ref](./goroutines.go)

## Memory Management

