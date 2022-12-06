`must`
=======

A library for _Must_ assertions in Go.

Snippets
--------
### Go get
```bash
go get github.com/mniak/must
```

### Import
```go
github.com/mniak/must
```

Motivation
----------
While developing scripts in Go, it can become very boring to handle each and every error.

In many places a function `Must` have appeared to overcome scenarios where you are sure that the error will never occur or you just want to stop execution when it happens.

That is even the case for some packages in the Go standard library. For instance:

- [Package `text/template`](https://pkg.go.dev/text/template#Must) has a function `Must`
- [Package `html/template`](https://pkg.go.dev/html/template#Must) has a function `Must`
- [Package `regexp`](https://pkg.go.dev/regexp#MustCompile) has a function `MustCompile`

I caught myself many times implementing `handleError` functions for command line applications and since Go 1.18 with generics, I've been creating in multiple repositories the same generic functions `Must0`, `Must1`, etc. Therefore I thought it would be better to create a library that could be used in various projects without duplicating the code.


Usage
-----

There are functions named `DoX` where `X` is the number of arguments except the error argument.

If the error is not nil, the _fatal_ function will be called.
The default function is `DefaultFatal` but you can set it using `SetFatalFunc`.

### Examples

#### Calling function that returns 1 value argument and an error
```go
package main

import (
	"errors"
	"fmt"

	"github.com/mniak/must"
)

func myFunction(n int) (string, error) {
	return "", errors.New("a fatal error has occurred")
}

func main() {
	value := must.Do1(myFunction(123))
	fmt.Println(value)
}

```

#### Setting a custom _fatal_ function
```go
package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/mniak/must"
)

func myFunction(n int) (string, error) {
	return "", errors.New("a fatal error has occurred")
}

func main() {
	must.SetFatalFunc(func(err error) {
		log.Fatalf("A fatal error occurred: %s", err.Error())
	})
	value := must.Do1(myFunction(123))
	fmt.Println(value)
}

```

Default _fatal_ behavior
------------------------
The default fatal behavior is to call the function `DefaultFatal` which in turn writes the error message into `stderr` followed by a _new line character_ and exits with status `-1 (255)`.

