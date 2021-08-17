# GoLang

## Preface

This documentation mainly consists of materials found at

* https://golang.org/doc/tutorial/getting-started
* https://tour.golang.org

with my own examples in the repository.

## Setup and general commands

First we need to create a module.
Modules are the units of distribution and versioning.
We can create a module by running the following command

```sh
go mod init <name of the module>
```

Typically the name of the module is the location of the Internet where you would install that module.
For example

```sh
# go mod init github.com/moledoc/letsgo
go mod init example.com/user/example
```

where `go.mod` is

```go
module example.com/user/example

go <go version>
```

Module name can look like `github.com/<git username>/<repo name>`.

This command creates a file named `go.mod`.

Next we will create a `go` package: it is a directory containing `go` code.
So, for example (will write as commands, but logic applies to IDE's/editors, such as vscode as well):

```sh
mkdir example
cd example
touch example.go
```

where `example.go` is

```go
package example

// Example returns the sentence 'This is an example'
func Example() string{
	return "This is an example"
}

```

Next we need a main function, that is located in the root directory of the module.
So, for example, when we are in the directory example/, then:

```sh
cd ..
touch main.go
```

where `main.go` is

```go
package main

import (
	"fmt"

	example "example.com/user/example.go"
)

func main(){
	fmt.Println(example.Example())
}
```

We can 

* run this `go` program by running the command 

```sh
go run main.go
```

* build this `go` program by running the command (creates binary `main` to the current directory)

```sh
go build main.go
```


* install this `go` program by running one of the commands (creates binary `main` to the directory $HOME/go/bin)

```sh
go install example.com/user/example
go install .
go install
```

* uninstall this `go` program by running the command

```sh
go clean -i
```

To run installed binary, execute command

```sh
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
main
```

or add `$HOME/go/bin` to the `$PATH` and then the command `main` is sufficient.

To clean modcache, run

```sh
go clean -modcache
```

To use a package from the internet, we need to download it (simple in vscode) and then add the package name to `go.mod` file,
so that the module knows, that we are using this package.
For example

```go
module example.com/user/example

go 1.16

require github.com/google/go-cmp v0.5.6
```

## Testing

To write tests, we create a file with suffix `_test` to the package directory

```sh
touch example/example_test.go
```

where the contents is

```go
package example_test

import (
	"testing"

	example "example.com/user/example"
)

func TestExample(t *testing.T){
	if example.Example() != "This is an example" {
		t.Fatal("<Descriptive error message>")
	}
}
```

To run the test, navigate to the package directory and run

```sh
go test
```

When the test file contains multiple test, then we get more verbose output, when we run

```sh
# when inside package directory
go test -v
# or
# when in the root directory for all packages
go test -v ./...
# or
# when in the root directory for selected package
go test -v ./<package name>
```

When function name starts with capital letter, then that function is exported.
Functions with lower case names are not exported.

## Basic types

* bool
* string
* int,, int8,, int16,, int32,, int64, uint, uint8, uint16, uint32, uint64, uintptr
* byte // alias for uint8
* rune // alias for int32; represents a unicode code point
* float32, float64
* complex64, complex128

## Assigning variables

Assigning value to a variable

```go
assigning := "This is a way to value a variable"
var assigningVar = "This is another  way to value a variable, but declares a list of value and can have type as well"
var assigningVarV2 string
assigningVarV2 = "this assign has explicit type declared"
var fst,snd,thrd,frth = "first",1,true,2.0
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

Assigning value to variable with `var` can be done at package or function level; with `:=` only at function level.
Variables without initial value are given the corresponding _zero_ value:

* `0` for numeric
* `false` for boolean
* `""` (empty string) for string

Type conversion

```go
var i int = 32
iFloat := float32(i)
var iFloatUint uint = uint(iFloat)
```

You can declare constants with keyword `const`.
This keyword can be used at package and function level.
They can be boolean, character, string or numeric values.


An untyped constant takes the type needed by its context (eg by function return type).
Example from https://tour.golang.org/basics/16

```go
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

## Loops

Only `for` loop exist in GoLang.
A `for` loop has 3 components:

* (optional) init expression, that is executed before the loop
* condition expression, that is evaluated before every loop iteration
* (optional) post expression, that is executed after every loop iteration

Variables declared in the init expression are visible only in the `for` loop scope.
Example of a basic `for` loop

```go
for i:=0; i<10; i++ {
	fmt.Println(i)
}
```

We can omit the init and post expression, so a `for` loop would look like

```go
for ; i<10; {
	i += i
}
```

In this case, we can omit the extra semicolons (;), making the `for` loop equivalent to typical `while` loop

```go
for i<10 {
	i += i
}
```

To exit a break loop, we can use `break` command.
For example

```go
for i<10 {
	if i==5 {
		break
	}
	i += i
}
```


Leaving out the condition expression results in infinite loop.

## if-else statement

A basic `if` statement in Go looks like this

```go
var value = 1
if value = 0 {
	fmt.Println("value = 0")
}
```

We can add init expression to the `if` statement.
That init expression is executed before the condition expression in the `if` clause.
Furthermore, any variables declared in the init expression are in scope until the end of `if` statement.
For example

```go
if value := 1; value = 0 {
	fmt.Println("value = 0")
}
```

Any variables declared in `if` statement are valid in the corresponding `else` statement as well.

```go
if value := 1; value = 0 {
	fmt.Println("value = 0")
} else {
	fmt.Sprintf("value = %v\n",value)
}
```

### switch statement

A `switch` statement is a shorter/more concise version of writing a `if-else` block.
In GoLang, the switch statement runs the first case whose value equals to the condition expression and not the cases that follows.
Basically `break` statement is automatically provided in GoLang.

In GoLang, switch cases need not be constants and the values involved need not be integers.

Example of a `switch` statement.

```go
switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("OS X.")
case "linux":
	fmt.Println("Linux.")
default:
	// freebsd, openbsd,
	// plan9, windows...
	fmt.Printf("%s.\n", os)
}
```

Also, case statement can have a function call (that function needs to return correct type for the case).
Without condition equals to `switch true` statement.
For example

```go
t := time.Now()
switch {
case t.Hour() < 12:
	fmt.Println("Good morning!")
case t.Hour() < 17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")
}
```

The `switch true` statement can be used to construct long `if-else` statements in a clean manner.


## defer

A `defer` statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

For example 

```go
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

returns `hello world`.

Deferred function calls can be stacked.
In that case, the deferred function calls are pushed onto a stack.
When a function returns, its deferred calls are executed in last-in-first-out order.

## Pointers

A pointer holds the memory address of a value and it can be *dereferenced* and *inderected* (in Go's language).
In Go there are no pointer arithmetic, unlike in C.
A pointers zero value is `nil`.

The `*` operator denotes the pointer's underlying value (dereferencing).
The `&` operator generates a pointer to its operand (inderecting).

For better understanding, look at the following example:

```go
i, j := 42, 2701

p := &i         // point to i
fmt.Println(*p) // read i through the pointer
*p = 21         // set i through the pointer
fmt.Println(i)  // see the new value of i

p = &j         // point to j
*p = *p / 37   // divide j through the pointer
fmt.Println(j) // see the new value of j

// output:
// 42
// 21
// 73
```

## Struct

A `struct` is a collection of fields.

For example

```go
package main

import "fmt"

type Example struct {
	X int
	Y int
}

func main() {
	fmt.Println(Example{1,2})
}
```

The fields of a `struct` are accessed using a dot.

```go
example := Example{1,2}
fmt.Println(example.X)
```

The fields of a `struct` can be accessed through a struct pointer.
Following previous examples, when we have a struct pointer `p` of struct `Example`, then we could access field `X` with `(*p).X`.
Since this notation is deemed cumbersome, the language permits writing `p.X`, without the explicit dereference.
Either one could be used.

```go
example := Example{1, 2}
p := &example
p.X = 1e9
fmt.Println(example)
(*p).X = 2e9
fmt.Println(example)
```


## Author

Written by
Meelis Utt
