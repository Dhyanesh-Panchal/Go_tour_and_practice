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

## Modules
Modern go uses Modules to manage packages.
- Creating a Module: ``go mod init modulename``
- Adding external packages: ``go get github.com/link_to/your_package``
- Update packages: ``go mod tidy``

