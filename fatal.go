package must

import (
	"fmt"
	"os"
)

var fatal func(err error) = DefaultFatal

func SetFatalFunc(fn func(err error)) {
	fatal = fn
}

func ResetFatalFunc() {
	fatal = DefaultFatal
}

func DefaultFatal(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}
}
