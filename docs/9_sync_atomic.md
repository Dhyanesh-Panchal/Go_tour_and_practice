# Synchronization of Goroutines using `sync/atomic`

The `sync/atomic` package in Go provides low-level atomic memory primitives for synchronization.

## What is an Atomic Operation?
An operation is atomic if it completes in a single step relative to other goroutines. This means if one goroutine performs an atomic operation, no other goroutine can see the variable in an intermediate state.

## Functions in `sync/atomic`
The package provides atomic operations for various data types. Below is a categorized list of functions.

### Integer Functions (`int32`, `uint32`, `int64`, `uint64`)

#### `int32` Functions:
- `AddInt32(addr *int32, delta int32) int32` - Atomically adds `delta` to the value at `addr` and returns the new value.
- `CompareAndSwapInt32(addr *int32, old, new int32) bool` - Atomically swaps the value at `addr` with `new` if the current value equals `old`.
- `LoadInt32(addr *int32) int32` - Atomically loads and returns the value at `addr`.
- `StoreInt32(addr *int32, val int32)` - Atomically stores `val` at `addr`.
- `SwapInt32(addr *int32, new int32) int32` - Atomically swaps the value at `addr` with `new`, returning the old value.

#### `uint32` Functions:
- `AddUint32(addr *uint32, delta uint32) uint32`
- `CompareAndSwapUint32(addr *uint32, old, new uint32) bool`
- `LoadUint32(addr *uint32) uint32`
- `StoreUint32(addr *uint32, val uint32)`
- `SwapUint32(addr *uint32, new uint32) uint32`

**Similar functions exist for `int64` and `uint64`.**

### Functions for `uintptr`
- `AddUintptr(addr *uintptr, delta uintptr) uintptr`
- `CompareAndSwapUintptr(addr *uintptr, old, new uintptr) bool`
- `LoadUintptr(addr *uintptr) uintptr`
- `StoreUintptr(addr *uintptr, val uintptr)`
- `SwapUintptr(addr *uintptr, new uintptr) uintptr`

### Functions for Pointers (`unsafe.Pointer`)
- `CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) bool`
- `LoadPointer(addr *unsafe.Pointer) unsafe.Pointer`
- `StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)`
- `SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) unsafe.Pointer`
