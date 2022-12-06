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
