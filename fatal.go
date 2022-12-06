package must

import (
	"log"
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
		log.Fatal(err)
	}
}
