# GoLang

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

To use a package from the internet, we need to download it (simple in vscode) and then add the package name to `go.mod` file,
so that the module knows, that we are using this package.
For example

```go
module example.com/user/example

go 1.16

require github.com/google/go-cmp v0.5.6
```

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

## Author

Written by
Meelis Utt
