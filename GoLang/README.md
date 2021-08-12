# GoLang

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
go test -v
# or
go test -v ./...
# or
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

* (optional) init statement, that is executed before the loop
* condition statement, that is evaluated before every loop iteration
* (optional) post statement, that is executed after every loop iteration

Variables declared in the init statement are visible only in the `for` loop scope.
Example of a basic `for` loop

```go
for i:=0; i<10; i++ {
	fmt.Println(i)
}
```

We can omit the init and post statement, so a `for` loop would look like

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

Leaving out the conditional statement results in infinite loop.

## if-else statement

A basic `if` statement in Go looks like this

```go
var value = 1
if value = 0 {
	fmt.Println("value = 0")
}
```

We can add init statement to the `if` statement.
That init statement is executed before the conditional statement in the `if` clause.
Furthermore, any variables declared in the init statement are in scope until the end of `if` statement.
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

## Author

Written by
Meelis Utt
