# Golang Playground

This repository holds [Go](https://go.dev/) snippets created while I'm recalling and learning the language

## References

* [Effective Go](https://go.dev/doc/effective_go)
* [A Tour of Go](https://go.dev/tour/)
* [Go Programming](https://www.youtube.com/watch?v=CF9S4QZuV30) ([Derek Banas](https://www.youtube.com/@derekbanas))
  * [Cheat sheet](https://www.newthinktank.com/2015/02/go-programming-tutorial/)

## Packages vs Modules

From [How to Write Go Code](https://go.dev/doc/code):

* **Package:**
  * Go programs are organized into packages.
  * A package is a collection of source files in the same directory that are compiled together.
  * Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.
* **Module:**
  * A repository contains one or more modules.
  * A module is a collection of related Go packages that are released together.
  * A Go repository typically contains only one module, located at the root of the repository.
  * A file named go.mod there declares the module path: the import path prefix for all packages within the module

Adapted example from [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module):

```
$ pwd
<HOME>/go/src/example.com

$ mkdir greetings

$ cd greetings

$ go mod init example.com/greetings
go: creating new go.mod: module example.com/greetings

$ ls
go.mod

$ cat go.mod
module example.com/greetings
go 1.21.2

$ vi greetings.go

$ cat greetings.go
package greetings
import "fmt"
func Hello(name string) string {
    return fmt.Sprintf("Hi, %v. Welcome!", name)
}

$ vi greetings_test.go 

$ cat greetings_test.go
package greetings
import (
    "fmt"
    "testing"
)
func TestHello(t *testing.T) {
    got := Hello("World")
    want := "Hi, World. Welcome!"
    if got != want {
        t.Errorf("got != want | '%s' != '%s'", got, want)
    }
    fmt.Printf("ok: got == want == '%s'\n", got)
}

$ go test
ok: got == want == 'Hi, World. Welcome!'
PASS
ok      example.com/greetings   0.002s
```

## Core

- [main.go](main.go): executable, command-line parsing, [type assertions (any / interface{})](https://go.dev/tour/methods/15) ...
- [c_call.go](c_call/c_call.go): call C code from Go
  - [Calling C code from go](https://karthikkaranth.me/blog/calling-c-code-from-go/)
  - [Call C code from Golang](https://medium.com/@vivek2793/call-c-code-from-golang-8783c6b58a5c)
- [concurrency.go](concurrency/concurrency.go): goroutines, sync, channels, context, timeouts, select, ...
  - [Advanced Golang: Channels, Context and Interfaces Explained - Code With Ryan](https://www.youtube.com/watch?v=VkGQFFl66X4)
  - [Concurrency in Go - Jake Wright](https://www.youtube.com/watch?v=LvgVSSpwND8)
  - [how to listen to N channels? (dynamic select statement)](https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement)
  - [Go → Reflect → SelectCase](https://pkg.go.dev/reflect#SelectCase)
- [ctx.go](ctx/ctx.go): simple context example
  - [Go Concurrency Patterns: Context](https://go.dev/blog/context)
  - [Context package godoc](https://pkg.go.dev/context): Context type carries deadlines, cancellation signals and other request-scoped values across API boundaries and between processes.
- [http_json.go](http_json/http_json.go): HTTP Server, HTTP Client, JSON, ...
- [make_vs_new.go](make_vs_new/make_vs_new.go): make vs new
  - [Why would I make() or new()?](https://stackoverflow.com/questions/9320862/why-would-i-make-or-new)
  - [Golang New vs Make](https://medium.com/learn-code/golang-new-vs-make-8a4dbd84e92b)
- [pointers_refs.go](pointers_refs/pointers_refs.go): pointers and references
  - [Should I define methods on values or pointers?](https://go.dev/doc/faq#methods_on_values_or_pointers):
    - Does the method need to modify the receiver?
    - It will be much cheaper to use a pointer receiver.
    - If some of the methods of the type must have pointer receivers, the rest should too, so the method set is consistent regardless of how the type is used.
  - [Why do T and *T have different method sets?](https://go.dev/doc/faq#different_method_sets)
    - The method set of a type T consists of all methods with receiver type T
    - That of the corresponding pointer type *T consists of all methods with receiver *T or T.
    - That means the method set of *T includes that of T, but not the reverse.
- [reflection.go](reflection/reflection.go): reflection
  - [Go (Golang) Reflection Tutorial](https://www.youtube.com/watch?v=f4aUrm7N5DU)
- [structs_ints.go](structs_ints/structs_ints.go): structs and interfaces

