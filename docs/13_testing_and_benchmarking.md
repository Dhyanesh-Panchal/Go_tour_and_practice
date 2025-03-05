# Testing in Go
Code Ref: [here](../Practice/go-testing)
In go, test files are decalred with filename "_test" in the suffix.
- The file will be excluded from regular package builds but will be included when the "go test" command is run.
- The test file can be in the same package as the one being tested, or in a corresponding package with the suffix "_test".
- If the test file is in the same package, it may refer to unexported identifiers within the package.
- If the file is in a separate "_test" package, the package being tested must be imported explicitly and only its **exported identifiers may be used**. This is known as "black box" testing.

We can run test using `go test`

- To calculate and export the coverage for our test, we can simply use the -coverprofile argument with the go test command.
```bash
go test ./maths -coverprofile=outputfile.out
```

## `testing` package
- Go provided inbuilt package for writing test codes for packages.
- We can use `testing.T` as a argument in testing funciton. Eg.
```go
package abs

import "testing"

func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```

### Errors and Logs
- The `testing.T` type provides various practical tools to interact with the test workflow, including `t.Errorf()`, which prints out an error message and sets the test as failed.
- It is important to mention that `t.Error*` **does not stop the execution of the test**.  Instead, all encountered errors will be reported once the test is completed.
- If we need to Stop the execution of further tests, in that case we can ise `t.Fatal`


### CleanUps
- Each type (`testing.T`,`testing.B`,`testing.F`,..etc) have `Cleanup()` funtionality.
- Cleanup registers a function to be called when the test (or subtest) and all its subtests complete. Cleanup functions will be called in last added, first called order. (very similar to defer)
Eg. of benchmarking:
```go
func setup() string {
	tempFile, _ := os.CreateTemp("", "benchmark")
	defer tempFile.Close()
	return tempFile.Name()
}

func cleanup(filename string) {
	os.Remove(filename)
}

func BenchmarkWithCleanup(b *testing.B) {
	file := setup()
	b.Cleanup(func() { cleanup(file) }) // Ensures cleanup runs after benchmark
	for i := 0; i < b.N; i++ {
		os.WriteFile(file, []byte("test"), 0644)
	}
}

```
## Table Driven Tests
- `t.Run()` can be used to run multiple subtests. 
- signature: ``func (t *T) Run(name string, f func(t *T)) bool``
- `Run` runs `f` as a subtest of `t` called name.
- It runs `f` in a separate goroutine and blocks **until f returns or calls `t.Parallel` to become a parallel test**.

Eg.
```go
type adderTestCase struct {
	testName       string
	a, b, expected int
}

var tests = []adderTestCase{
	{"Test Case: #1", 1, 2, 3},
	{"Test Case: #2", 3, 2, 5},
	{"Test Case: #3", 0, 4, 4},
	{"Test Case: #4", -5, 4, -1},
}

func TestAdder(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			got := maths.Adder(tc.a, tc.b)

			if got != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, got)
			}
		})

	}
}
```


## Fuzz testing
- With fuzzing, random data is run against your test in an attempt to find vulnerabilities or crash-causing inputs.

Rules to follow while writing fuzz test:
- A fuzz test must be a function named like _FuzzFuncname_, which accepts only a `*testing.F`, and has **no return value**.
- Fuzz tests must be in *_test.go files to run.
- A fuzz target must be a method call to `(*testing.F).Fuzz` which accepts a `*testing.T` as the first parameter, followed by the fuzzing arguments. There is no return value.
- There must be exactly one fuzz target per fuzz test.
- 
- 

- The unit test has limitations, namely that each input must be added to the test by the developer. One benefit of fuzzing is that it comes up with inputs for your code, and may identify edge cases that the test cases you came up with didn’t reach.
- Fuzzing has a few limitations as well. In your unit test, you could predict the expected output of the Reverse function, and verify that the actual output met those expectations.
- Example: [Fuzzy Testing with Reverse function](../Practice/go-testing/fuzzytesting)
- For example, in the test case Reverse("Hello, world") the unit test specifies the return as "dlrow ,olleH".
- When fuzzing, **you can’t predict the expected output**, since you don’t have control over the inputs.

## Benchmarking 
- benchmarking is done using `testing.B` type.
- It helps measure the execution time of functions by running them multiple times and computing an average time per operation.
- Benchmark functions must start with 'Benchmark', take a `*testing.B` argument, and reside in `_test.go` files.

**Note: It is possible to keep unit tests, benchmarks, and fuzz tests in the same `*_test.go` file**

