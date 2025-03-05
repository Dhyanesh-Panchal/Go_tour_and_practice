# Panics
Panics in Go are runtime errors that cause a program to abruptly halt execution.

Panic is a built-in function that stops the ordinary flow of control and begins panicking.

When the function `F` calls `panic`, execution of `F` stops, any deferred functions in `F` are executed normally, and **then F returns to its caller. To the caller, F then behaves like a call to panic.** 

The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes.

