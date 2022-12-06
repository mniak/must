package main

import (
	"fmt"
	"os"
)

func handle(err error) {
	if err != nil {
		fatal(-1, err.Error())
	}
}

func fatal(code int, msg string, msgargs ...any) {
	fmt.Fprint(os.Stderr, "Error: ")
	fmt.Fprintf(os.Stderr, msg, msgargs...)
	fmt.Fprintln(os.Stderr)
	os.Exit(code)
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: go run ./internal/cmd/gen MAX")
	fmt.Fprintln(os.Stderr, "  MAX: represents the maximum number of type arguments for which the functions will be generated.")
	fmt.Fprintln(os.Stderr)
}
