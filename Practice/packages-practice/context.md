# `context`
Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries to all the goroutines involved in handling a request.
- It helps prevent resource leaks by canceling operations that exceed a time limit.
- It propagates cancellation signals to multiple goroutines.
- It allows passing key-value pairs (Request Scoped Data) within a request's execution scope.
-  When a Context is canceled, all Contexts derived from it are also canceled.

The `WithCancel`, `WithDeadline`, and `WithTimeout` functions take a Context (the parent) and return a derived Context (the child) and a CancelFunc.
### Ways to create a new context:
- `context.Background()`: it returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
- `context.TODO()`: TODO returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available.
- `context.WithValue()`: 


