package must

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestDo0(t *testing.T) {
	t.Run("When does not have error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		Do0(nil)

		assert.Zero(t, fatalCount, "fatal must have never been called")
		assert.Nil(t, fatalErr)
	})

	t.Run("When has error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeErr := errors.New(gofakeit.SentenceSimple())
		Do0(fakeErr)

		assert.Equal(t, 1, fatalCount, "fatal must have called the fatal func exactly once")
		assert.Equal(t, fakeErr, fatalErr)
	})
}
