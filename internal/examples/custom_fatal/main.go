package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mniak/must"
)

func myFunction(n int) (string, error) {
	return "", errors.New("a fatal error has occurred")
}

func main() {
	must.SetFatalFunc(func(err error) {
		fmt.Println("Bye bye: graceful stop")
		os.Exit(0)
	})
	value := must.Do1(myFunction(123))
	fmt.Println(value)
}
