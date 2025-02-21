# Packages
Every Go program is made up of packages.
Programs start running in package `main`.

By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package comprises files that begin with the statement package `rand`.


### Types of package
1. **Executable Package**: These contain a main package and a main() function.
   They generate an executable binary.
2. **Library Packages**: They don't have `main` program, rather they are imported in other programs.
```go
package myutility

import "fmt"

func init() {
	fmt.Println("Initializing utils package...")
}

func Add(a, b int) int {
	return a + b
}
```
- `main` is a special package which indicates that it contains code that *can be built into binary and executed*.
- `main` package must contain `main()` function else it will give compile time error.
- Also, if we try to execute/build a Non-main package, it will thrown a compile time error.
#### `init()` function
- Each Packages defines `init()` function that runs before `main()`
- each file can have multiple init functions.
- `init()` is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized.
- If a package is imported by another package, the init functions of the imported package will be executed before the init functions of the importing package.
- Hence, in the example, the `init()` of `"fmt"` will be called first and then the init of this package.
- a common use of init functions is to verify or repair correctness of the program state before real execution begins.

For more info refer [here](./custom_packages/main.go)

### Imports
Used to import packages.
Its good idea to use *factored* import statment rather than using multiple imports. i.e. import (//all imports here)
- When importing a package from the standard library you need to use the full path to the package in the standard library tree, not just the name of the package. For example:
```go
import (
    "fmt"
    "math/rand"         // Not "rand"
    "net/http"          // Not "http"
    "net/http/httptest" // Not "httptest"
)
```

- If you import a package but don't actually use it in your code, it will result in a compile-time error.
### Exports
- Anything in Go code is exported if its name starts with a capital letter. Otherwise it is unexported.
- The difference between them is:
  - Unexported things are 'private' to the package that they are declared in. They are only visible to code in the same package.
  - In contrast, exported things in a package are 'public' and are visible to any code that imports the package, and can Modify it (unless they are `const`)
## Modules
Modern go uses Modules to manage packages.

A module is a tree of Go source files with a go.mod file in the tree's root directory.

- Creating a Module: ``go mod init modulename``
- Adding external packages: ``go get github.com/link_to/your_package``
- Update packages: ``go mod tidy``
- To list the reason, why particular module is present (in case of indirect dependencies): ``go mod why -m package/name/here``
   e.g:
   ```bash
   go mod why -m golang.org/x/sys
   # golang.org/x/sys
   lucky-number.alexedwards.net/cmd/cli
   github.com/fatih/color
   github.com/mattn/go-isatty
   golang.org/x/sys/unix
   ```

