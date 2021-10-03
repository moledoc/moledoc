# GoLang

## Preface

This documentation mainly consists of materials found at

* https://golang.org/doc/tutorial/getting-started
* https://tour.golang.org
* https://golang.org/doc/effective_go

with some of my own examples and some found at mentioned pages.

Some additional links are mentioned under subtopics, when necessary.

## Setup and general commands

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
* int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
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

**Type conversion**

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

To exit a break loop, we can use `break` command and to continue to next iteration, we can use `continue`
For example

```go
for i<10 {
	if i==5 {
		break
	}
	if i==4{
		i+=2
		continue
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

### Struct literal

A struct literal denotes a newly allocated struct value by listing the values of its fields.
You can list just a subset of fields by using the `Name:` syntax. (And the order of named fields is irrelevant.)
The special prefix `&` returns a pointer to the struct value.
For example

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
// output
// {1 2} {1 0} {0 0} &{1 2}
```

## Arrays

The type `[n]T` is an array of `n` values of type `T`.
So for example, the variable `exampleVar` is an array of 10 strings.

```go
var exampleVar [10]string
```

Go's arrays cannot be resized.


## Slice

Slices are dynamically-sized, flexible views into the elements of an array.
The type `[]T` is a slice with elements of type T.
A slice is formed by specifyng two indices: lower and upper bound, where lower bound is **included** and upper bound is **excluded**.
For example, if we have an array `a`, then we can make a slice from it as

```go
a[lower : upper]
```

Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.

Slices can contain any type, including other slices.

### Slice literal

A slice literal is like an array literal without the length.
So for example, an array literal would be

```go
[3]bool{true,true,false}
```

and corresponding slice literal would be

```go
[]bool{true,true,false}
```

Notice, that slice literal creates the above array and then references it.

### Slice defaults

When slicing, the lower and/or upper bound can be omitted.
In that case the default values are used.
The default value for lower bound is 0 and for upper bound is the length of the slice.
For example, when we have the following array, then all presented slices are equivalent.

```go
var a = [10]int

a[0:10]
a[0:]
a[:10]
a[:]
```

### Slice length and capacity

Length of the slice is the number of element the slice contains.
Capacity is the number of elements in the underlying array, counting from the first element in the slice.
Slice's length can be extended by re-slicing, but only when there is suffecient capacity.

We can find length and capacity of slice `s` with

```go
len(s)
cap(s)
```

### Nil slices

The zero value of a slice is `nil`.
In that case the length and capacity is 0 and the slice has no underlying array.

### Dynamic slices

Dynamic slices can be made by using built-in `make` function.
The `make` function allocates a zeroed array and returns a slice that refers to that array.
In the function `make`, the capacity of the slice can also specified.

```go
// capacity not specified
a := make([]int, 3)
// capacity specified
a := make([]int, 3,17)
```

### Appending to a slice

In GoLang there is a built-in `append` function.

```go
func append(s []T, vs ...T) []T
```

If the backing array is to small to fit all provided elements, then a bigger array will be allocated and returned slice will point to that array.

## Range

The `range` form of the `for` loop iterates over a slice or map.
When iterating over a slice, two values are returned each iteration: index, copy of the element at that index.

Example code

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
// output:
// 2**0 = 1
// 2**1 = 2
// 2**2 = 4
// 2**3 = 8
// 2**4 = 16
// 2**5 = 32
// 2**6 = 64
// 2**7 = 128
```

To skip index of value we can assign either one to `_`.
If we only want the index, we can omit the value alltogether.

```go
# skipping value
for i, _ := range pow
# skipping index
for _, value := range pow
# only index
for i := range pow
```


## Maps

Map maps keys to values.
We can initialize a map with function `make`.

```go
m := make(map[string]int)
```

The zero value of a map is `nil`: it has no keys and no keys can be added to it.

Map **literals** are analogous to struct literals, but the keys are required.

```go
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	}
}
```

If the top-level is just a type name, then we can omit it.

```go
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```

### Mutating maps

* Insert/update element: `m[key] = elem`
* Retrieve element: `elem = m[key]`
* Delete an element: `delete(m,key)`
* Test that key is present in map (with two value assignment): `elem, ok = m[key]`
  * if `key` is in `m`, then `ok` is `true` and `elem` is the element, else `ok` is `false` and `elem` is the zero value of the corresponding type.
  * if `elem`, `ok` are not declared yet, we can use shorthand declaration: `elem, ok := m[key]`.


## Function values

In GoLang functions are values as well (like in haskell, R etc (functional paradigm)).
They can be used as function arguments or return values.

### Function closures

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the `adder` function returns a closure.
Each closure is bound to its own `sum` variable.

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
//output
// 0 0
// 1 -2
// 3 -6
// 6 -12
// 10 -20
// 15 -30
// 21 -42
// 28 -56
// 36 -72
// 45 -90
```

## Methods

GoLang does not have classes, but we can define methods on types.
A method is a **function** with special _receiver_ argument.
The receiver argurment appears between keyword `func` and method name.

For example

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
// output:
// 5
```

Method can be declared with non-struct types as well.

```go
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
```

You can **only** declare a method with a receiver whose type is defined in the same package as the method.
You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).

You can declare methods with pointer receivers.
That means we have literal syntax `*T` as the receiver type (`T` can't be pointer itself).
Methods with pointer receivers can modify the value to which the receiver points.
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

We can write the `Scale` method as a function as well (`ScaleFun`).

Example

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleNoPtr(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v.Abs())
}

func ScaleFun(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.ScaleNoPtr(10)
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
	ScaleFun(&v, 10)
	fmt.Println(v.Abs())
}

// output:
// 50
// 5
// 50
// 500
```

### Methods and pointer indirection

Functions with pointer argument **must** take a pointer value.

```go
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

However, methods with pointer receivers take either a value or pointer as the receiver when they are called.

```go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

Since the `Scale` method has a pointer receiver,
then Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)`.

The equivalent happens in the reverse direction.
Functions that take a value argument must take a value of that specific type:

```go
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

while methods with value receivers take either a value or a pointer as the receiver when they are called:

```go
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

Here `p.Abs()` is interpreted as `(*p).Abs()`.

### Choosing value or a pointer receiver

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both `Scale` and `Abs` are with receiver type `*Vertex`, even though the `Abs` method needn't modify its receiver.

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

// output
// Before scaling: &{X:3 Y:4}, Abs: 5
// After scaling: &{X:15 Y:20}, Abs: 25
```

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.


## Interfaces

An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.

```go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	fmt.Println(a.Abs())

	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// output
// 1.4142135623730951
// 5
```

### Interfaces are implemented implicitly

A type implements an interface by implementing its methods.
There is no explicit declaration of intent, no `implements` keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

// output:
// hello
```

### Interface values

Interface values can be thought as a tuple of values `(value, type)`.
An interface value holds a value of a specific underlying concrete type.
Calling a method on an interface value executes the method of the same name on its underlying type.

```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// output:
// (&{Hello}, *main.T)
// Hello
// (3.141592653589793, main.F)
// 3.141592653589793
```

### Interface values with nil underlying values

If the concrete value inside the interface itself is `nil`, the method will be called with a nil receiver.
In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method `M` in this example.)

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// output:
// (<nil>, *main.T)
// <nil>
// (&{hello}, *main.T)
// hello
```

**Note** that an interface value that holds a nil concrete value is itself non-nil.

### Nil interface values

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

```go
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// output:
// (<nil>, <nil>)
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x47f427]
```

### Empty interface

The empty interface type specifies zero methods `interface{}`.
An empty interface may hold values of any type, since every type implements at least zero methods.
Empty interfaces are used by code that handles values of unknown type.
Good example is `fmt.Print` that thakes any number of arguements of type `interface{}`.

```go
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// output
// (<nil>, <nil>)
// (42, int)
// (hello, string)
```

### Type assertion

Type assertion provides access to an interface value's underlying concrete value

```go
t := i.(T)
```

This statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.
If `i` does not hold a `T`, then the statement will trigger a panic.

Type assertion can return two values (underlying value and a boolean value), so we can **test** whether an interface holds a specific type.

```go
t, ok := i.(T)
```

If `i` holds a `T`, then `t` will be the underlying value and `ok` will be `true`.
If not, then `t` will be the type `T` zero value and `ok` will be `false`.
This will not trigger a panic.
This syntax is similar to reading from a map.

```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

// output:
// hello
// hello true
// 0 false
// panic: interface conversion: interface {} is string, not float64
```

### Type switches

A type switch is a construct that permits several type assertions in series.
It is like a regular `switch` statement, but instead of values we specify types against the given interface value's type.

```go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

The declaration in a `type switch` has the same syntax as type assertion, but the specific type is replace with keyword `type`, so `i.(type)`.

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}

// output
// Twice 21 is 42
// "hello" is 5 bytes long
// I don't know about type bool!
```

### Stringers

One of the most ubiquitous interfaces is `Stringer` defined by the `fmt` package.

```go
type Stringer interface {
	String() string
}
```

A `Stringer` is a type that can describe itself as a string.
The `fmt` package (and others) look for this interface to print values.

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

// output
// Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
```

Another example

```go
package main

import "fmt"

type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v",addr[0],addr[1],addr[2],addr[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// output w/o Stringer func
// loopback: [127 0 0 1]
// googleDNS: [8 8 8 8]

// output w/ Stringer func
// loopback: 127.0.0.1
// googleDNS: 8.8.8.8
```

## Errors

In GoLang errors are expressed with `error` values.
The `error` type is a built-in interface similar to `fmt.Stringer`.

```go
type error interface {
	Error() string
}
```

When printing values, `fmt` package will (besides `fmt.Stringer`) look for `error` interface.
In GoLang the calling code should handle errors by testing wheter the error equals `nil`, where a nil `error` denotes success and non-nil `error` denotes failure.

A simple example

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

More comprehensive example of `error` interface

```go
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// output
// at <timestamp>, it didn't work
```

## Readers

The `io` package specifies the `io.Reader` interface, which represents read end of a stream of data.
The Go standard library contains [many implementations](https://cs.opensource.google/search?q=Read%5C(%5Cw%2B%5Cs%5C%5B%5C%5Dbyte%5C)&ss=go%2Fgo) of this interface, including files, network connections, compressors, ciphers, and others.

This interface (`io.Reader`) has a `Read` method

```go
func (T) Read(b []byte) (n int, err error)
```

`Read` populates given byte slice with data and returns the number of bytes populated and an error value.
When a stream ends, an `io.EOF` error is returned.

The following example code creates a `strings.Reader` and consumes its output 8 bytes at a time.

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

//output
// n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
// b[:n] = "Hello, R"
// n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
// b[:n] = "eader!"
// n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
// b[:n] = ""
```


## Images

The package `image` defines `Image` interface.

```go
package image
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

**Note**: the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.

The documentation of package [image](https://golang.org/pkg/image/#Image).
The `color.Color` and `color.Model` types are also interfaces,
but we'll ignore that by using the predefined implementations `color.RGBA` and `color.RGBAModel`.
These interfaces and types are specified by the `image/color` package.

```go
package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

// output
// (0,0)-(100,100)
// 0 0 0 0
```

## Goroutines

A **goroutine** is a lightweight thread managed by the Go runtime.

To start a goroutine with function `f(x,y,z)` we write

```go
go f(x,y,z)
```

The evaluation of `f`, `x`, `y` and `z` happens in the current goroutine, but the executon of `f` happens in the new goroutine.

Since goroutines run in the same address space, then the access to shared memeory must be syncronized.
The `sync` package provides useful primitives, but the are not as often needed, because Go has more other primitives.

Example of a simple program with goroutine

```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%v: %v\n",i,s)
	}
}

func main() {
	go say("world")
	say("hello")
}

// output
// 0: world
// 0: hello
// 1: world
// 1: hello
// 2: hello
// 2: world
// 3: world
// 3: hello
// 4: hello
```

### Channels

Channels are a typed conduit through which you can send and receive values with the channel operator `<-`

```go
ch <- v // Send v to channel ch
v:= <-ch // Receive from ch and assign value to v.
```

Analog to maps and slices, channels must be created before use

```go
ch := make(chan int)
```

By default, sends and receives block until other side is ready.
This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

// output
// -5 17 12
```

### Buffered Channels

In GoLang the channels can be _buffered_.
To buffer a channel, we can give buffer length as the second argument to `make` to initialize a buffered channel.

```go
ch := make(chan int, 100)
```

Sends to buffered channel block only then the buffer is full.
Receives block when the buffer is empty.

```go
package main

import "fmt"

func sum(s []int, c chan int,i int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
	fmt.Printf("Routine %v sent to channel\n",i)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0,5,-1,15}

	c := make(chan int,2)
	go sum(s[2*len(s)/3:], c,1)
	go sum(s[len(s)/3:2*len(s)/3], c,2)
	go sum(s[:len(s)/3], c,3)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

// possible outcome
// Routine 3 sent to channel
// 17
// Routine 1 sent to channel
// 19
// Routine 2 sent to channel
// -5
```

### Range and Close

A sender can `close` a channel to indicate that no more values will be sent.
Receivers can test whether a channel has been closed by assigning second parameter to the receive expression, similar to key in map check.

```go
v, ok := <-ch
```

The parameter `ok` is `false` if there are no more values to receive and the channel is closed.

The loop `for i := range c` receives values from the channel repeatedly until it is closed.

**NOTE**: only the sender should close a channel, never the receiver.
Sending on a closed channel will casue a panic.
**NOTE**: Channels differ from file, because channels usually are closed for you. Closing is only necessary when the receiver must be told that there are no more values coming, eg to terminate `range` loop.

```go
package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		time.Sleep(200 * time.Millisecond)
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// output
// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
```

### Select

The `select` statement lets goroutine wait on multiple commumication operation.
A `select` blocks until on of its cases can run and then executes that case.
If multiple are ready at the same time, random case is selected.

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

We can use `default` case in a `select` statement.
The `default` case is run if no other case is ready.
It can be used to try a send or receive without blocking.

```go
select {
    case i := <-c
        // use i
    default:
        //receiveing from c would block
}
```

An example code

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// output
//     .
//     .
// tick.
//     .
//     .
// tick.
//     .
//     .
// tick.
//     .
//     .
// tick.
//     .
//     .
// tick.
// BOOM!
```

Example exercise

```go
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.

// Reason why function Walk is not recursive, is because then we can use the `close` function to close the channel
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkRecursive((*t).Left, ch)
	ch <- t.Value
	walkRecursive((*t).Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1Chan := make(chan int, 10)
	t2Chan := make(chan int, 10)
	go Walk(t1, t1Chan)
	go Walk(t2, t2Chan)
	for {
		t1Elem, ok1 := <-t1Chan
		t2Elem, ok2 := <-t2Chan
		// if elements differ, trees are not the same
		// if one tree ends before other, trees are not the same
		if t1Elem != t2Elem || ok1 != ok2 {
			return false
		}
		// if channels are closed, then break the loop
		// we can check only one channel, because we checked in prev if they are equal (ok1==ok2).
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	treeWalk := tree.New(1)
	fmt.Println("Test Walk function:")
	fmt.Println(treeWalk)
	go Walk(treeWalk, ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Test if trees are same:")
	fmt.Printf("\tIs the tree1 and tree1 same: %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("\tIs the tree1 and tree2 same: %v\n", Same(tree.New(1), tree.New(2)))
}

// output
// Test Walk function:
// ((((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10)
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
// Test if trees are same:
// 	Is the tree1 and tree1 same: true
// 	Is the tree1 and tree2 same: false
```

### sync.Mutex

We've seen how channels are great for communication among goroutines.

But what if we don't need communication?
What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?

This concept is called mutual exclusion, and the conventional name for the data structure that provides it is `mutex`.

Go's standard library provides mutual exclusion with `sync.Mutex` ans its methods `Lock` and `Unlock`.
We can define a block of code to be executed in mutual exclusion by surrounding it with a call to `Lock` and `Unlock` (see `Inc` method).

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

// output
// 1000
```

We can use `defer` to ensure that mutex will be unlocked as in the `Value` method.

### WaitGroup

* https://gobyexample.com/waitgroups

If we want to wait for multiple goroutines, we can use a **wait group**

Using `WaitGroup` is more efficient compared to using `sleep` to check if all goroutines are finished.
When we launch a goroutine,
we increment the `WaitGroup` counter with function `Add` and when a goroutine finishes, we decrement the counter with function `Done`.
If a `WaitGroup` is explicitly passed into function, then it should done by _pointer_.

```go
// To wait for multiple goroutines to finish, we can
// use a *wait group*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// This WaitGroup is used to wait for all the
	// goroutines launched here to finish. Note: if a WaitGroup is
	// explicitly passed into functions, it should be done *by pointer*.
	// This would be important if, for example, our worker had to launch
	// additional goroutines.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup
	// counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// Avoid re-use of the same `i` value in each goroutine closure.
		// See [the FAQ](https://golang.org/doc/faq#closures_and_goroutines)
		// for more details.
		i := i

		// Wrap the worker call in a closure that makes sure to tell
		// the WaitGroup that this worker is done. This way the worker
		// itself does not have to be aware of the concurrency primitives
		// involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Block until the WaitGroup counter goes back to 0;
	// all the workers notified they're done.
	wg.Wait()

	// Note that this approach has no straightforward way
	// to propagate errors from workers. For more
	// advanced use cases, consider using the
	// [errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup).
}

// possible output:
// worker 4 starting
// Worker 2 starting
// Worker 1 starting
// Worker 3 starting
// Worker 5 starting
// Worker 3 done
// Worker 5 done
// Worker 2 done
// Worker 4 done
// Worker 1 done
```

## Formatting

In Go formatting is easy.
It is left to `gofmt` program, that formats the go code according to set rules.
The program `gofmt` outputs the formatted code to stdout.

Some other notes about formating:

* tabs over spaces, but the latter can be used if needed.
* line length has no limit. for readability wraping and extra indent can be used.

## Commentary

* https://golang.org/doc/effective_go#commentary

Go provides C-style `/* */` block comments and C++-style `//` line comments.
_Line comments_ are the norm;
block comments appear mostly as package comments, but are useful within an expression or to disable a block of code.

We can use program `godoc` to process the Go source files and extract documentation about the contents of the package.
Comments that appear before top-level declarations, with no intervening newlines, are extracted along with the declaration to serve as explanatory text for the item.
Documentation produced with `godoc` is only good when the comments are good.

Every package should have a package comment, a block comment preceding the package clause.
For multi-file packages, the package comment only needs to be present in one file, and any one will do.
The package comment should introduce the package and provide information relevant to the package as a whole.
It will appear first on the `godoc` page and should set up the detailed documentation that follows.

Example of a package comment

```go
/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp
```

If the package is simple, we can use line comments to make a brief package comments.

```go
// Package path implements utility routines for
// manipulating slash-separated filename paths.
```

THe formatting of the comments are handled by `gofmt`, so no extra banners are needed.

The comments are uninterpreted plain text, so HTML and annotations such as `_this_` should _not_ be used.
One adjustment `godoc` does do is to display indented texxt in a fixed-width font, suitable for program snippets. Eg see [`fmt` package](https://golang.org/pkg/fmt/)

Depending on the context, `godoc` might not even reformat comments, so comments should look good: use correct spelling, punctuation and sentence structure, wrap lines etc.

Inside a package, any comment immediately preceding a top level declaration serves as a _doc comment_ for that declaration.
Every exported (capitalized) name in a program should have a doc comment.
Doc comments work best as complete sentences, which allow a wide variety of automated presentations.
The first sentence should be a one-sentence summary that starts with the name being declared.

```go
// Compile parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func Compile(str string) (*Regexp, error) {
```

If every doc comment begins with the name of the item it describes,
then using go [doc](https://pkg.go.dev/cmd/go#hdr-Show_documentation_for_package_or_symbol) subcommand and run the output through `grep`.
So for example, you have forgot the function name that parses regular expressions, you can find out the function name by running

```sh
go doc -all regexp | grep -i parse
```

But if all the doc comments in the package start with "This function ...",
then `grep` would be of little help finding the function name.

Go's declaration syntax allows grouping of declarations.
A single doc comment can introduce a group related constants or variables.
Since a whole declaration is presented, such a comment can often be perfunctory.

```go
// Error codes returned by failures to parse an expression.
var (
    ErrInternal      = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
    ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
    ...
)
```

Grouping can also indicate relationships between items, such as the fact that a set of variables is protected by a mutex

```go
var (
    countLock   sync.Mutex
    inputCount  uint32
    outputCount uint32
    errorCount  uint32
)
```

## Names

* https://golang.org/doc/effective_go#names

Naming is important (everywhere).
In go there is even a semantic effect: the visibility of a name outside a package is determined by whether the first character is upper case.

### Package names

After inmportig a package, the name of the package becomes an accessor for the content.
For example, after importing

```go
import "bytes"
```

the importing package can talk about `bytes.Buffer`.

A good package name is short, concise and evocative.
By convention, packages are given lower case, single-word names;
there should be no need for underscores or mixedCaps.
Err on the side of brevity, since everyone using the package will be typing that name.

There is no need to worry about package collisions.
The package name is only the default name for imports.
In case of collision, the importing package can choose a different name to use locally.

Another convention is that the package name is the base name of its source directory.
The package in `src/encoding/base64` is imported as "encoding/base64", but has name `base64`.

The importer of a package will use the name to refer to its coneents, so exported names in the package can use that fact to avoid repetition.
For instance, the buffered reader type in the `bufio` package is called `Reader`, not `BufReader`, because users see it as `bufio.Reader`, which is a clear and concise name.
Since imported entities are always addressed with their package name, then `bufio.Reader` and `io.Reader` do not conflict.
Another example is the function to make new instances of `ring.Ring`.
It is the definition of a _constructor_ in Go.
Normally it would be called `NewRing`, but since `Ring` is the only type exported by the package and since the package name is `ring`, then it's just called `New`.
So the user of the package just sees `ring.New`.
Use the package structure to help you choose good names.

Another short example is `once.Do`.
The `once.Do(setup)` reads well and would not be improved by writing `once.DoOrWaitUntilDone(setup)`.
Long names don't automatically make things more readable.
A helpful doc comment can often be more valuable than an extra long name.

### Getter

* https://golang.org/doc/effective_go#getter

In Go there are no automatic support for getters and setters.
It's ok to define own getters and setters, when it's appropriate to do so.
However, it's not idiomatic nor necessary to put `Get` into the getter's name.
If you have a field called `owner` (lower case, unexported), the getter method should be called `Owner` (upper case, exported), not `GetOwner`.
The use of upper-case names for export provides the hook to discriminate the field from the method.
A setter function, if needed, will likely be called `SetOwner`.
Both names read well in practice.

```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

### interface names

* https://golang.org/doc/effective_go#ginterface_names

By convention, one-method interfaces are named by the method name plus and -`er` suffix or similar modification to construct an agent noun: `Reader`, `Writer`, `Formatter`, `CloseNotifier` etc.

There are a number of such names and it's productive to honer them and the function names they capture.
`Read`, `Write`, `Close`, `Flush`, `String` and so on have canonical signatures and meanings.
To avoid confusion, don't give your methods one of those names, _unless_ they have the same signature and meaning.
Conversely, if your type implements a method with the same meaning as a method on a well-known type, give it the same name and signature: call ypourstring-converter method `String` and not `ToString`.

### MixedCaps

The convention in Go is to use `MixedCaps` or `mixedCaps` for multiword names.




## Author

Compiled by
Meelis Utt
