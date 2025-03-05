# `sync`
Package sync provides basic synchronization primitives such as mutual exclusion locks.

Under the sync package, following Types are available:

## WaitGroups
WaitGroups are use to manage go routines, and wait untill all the routine finishes.
It Provides three methods:
- Add(delta int) takes in an integer value which is essentially the number of goroutines that the WaitGroup has to wait for. This must be called before we execute a goroutine.
- Done() is called within the goroutine to signal that the goroutine has successfully executed.
- Wait() blocks the program until all the goroutines specified by Add() have invoked Done() from within.

## Mutex
A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
There are three funcitons:
- Lock(): Locks the Mutex, If the lock is already in use, the calling goroutine blocks until the mutex is available.
- Unlock(): Unlocks the mutex. **It is a run-time error if m is not locked on entry to Unlock.**
- TryLock(): TryLock tries to lock mutex and reports whether it succeeded.

## RWMutex
RWMutex is similar to Mutex, but it additionally provided functionality of Read-Locks. Additional methods:
- `RLock()` and `RUnlock()`
- This type doesn't have `TryLock()` method.

## Pool
- A Pool is a set of temporary objects that may be individually saved and retrieved.
- Any item stored in the Pool may be removed automatically at any time without notification. If the Pool holds the only reference when this happens, the item might be deallocated.
- A Pool is safe for use by multiple goroutines simultaneously.
- Pool's purpose is to cache allocated but unused items for later reuse, relieving pressure on the garbage collector.
- The appropriate use of a Pool is to manage a group of temporary items silently shared among and _potentially reused_ by concurrent independent clients of a package.
- It is important to note that Pool also has its performance cost. It is **much slower to use sync.Pool than simple initialization**.
It has two methods:
- `Get()` selects an arbitrary item from the Pool, removes it from the Pool, and returns it to the caller.
- `Put(x any)` adds the item to the pool.


## Map
It is like a Go `map[any]any` but is safe for concurrent use by multiple goroutines without additional locking or coordination.
- Loads, stores, and deletes run in amortized constant time.
Methods under sync.Map are:
- `Clear()` deletes all the entries, resulting in an empty Map.
- `CompareAndDelete(key, old any) (deleted bool)` deletes the entry for key if its value is equal to old. _The old value must be of a comparable type_.
- `Delete(key any)` Deletes the value of the key.
- `Load(key any) (value any,ok bool)` Load returns the value stored in the map for a key, or **nil if no value is present**.(Not the zero Value)
- `LoadAndDelete(key any) (value any,ok bool)` deletes the value for a key, returning the previous value if any. 
- `Range(f func(key,val any) bool)` This exectes f for key,value present in the map, **untill f returns false**.
- `Swap(key, val any) (previous any, loaded bool)` Swap swaps the value for a key and returns the previous value if any. The loaded result reports whether the key was present.
- `Store(key, value any)` sets the value for the key.
For code: [ref](./sync.go)

# Cond
- Cond implements a condition variable.
- Each Cond has an associated Locker L (often a *Mutex or *RWMutex), which must be held when changing the condition and when calling the `Cond.Wait` method.



# There is a subpackage `sync/atomic` which is used to provide lower level synchronization for primitive datatypes.






