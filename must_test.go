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

func TestDo1(t *testing.T) {
	t.Run("When does not have error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		value1 := Do1(fakeValue1, nil)

		assert.Zero(t, fatalCount, "fatal must have never been called")
		assert.Nil(t, fatalErr)
		assert.Equal(t, fakeValue1, value1)
	})

	t.Run("When has error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		fakeErr := errors.New(gofakeit.SentenceSimple())
		value1 := Do1(fakeValue1, fakeErr)

		assert.Equal(t, 1, fatalCount, "fatal must have called the fatal func exactly once")
		assert.Equal(t, fakeErr, fatalErr)
		assert.Equal(t, fakeValue1, value1)
	})
}

func TestDo2(t *testing.T) {
	t.Run("When does not have error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		fakeValue2 := gofakeit.SentenceSimple()
		value1, value2 := Do2(fakeValue1, fakeValue2, nil)

		assert.Zero(t, fatalCount, "fatal must have never been called")
		assert.Nil(t, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
	})

	t.Run("When has error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		fakeValue2 := gofakeit.SentenceSimple()
		fakeErr := errors.New(gofakeit.SentenceSimple())
		value1, value2 := Do2(fakeValue1, fakeValue2, fakeErr)

		assert.Equal(t, 1, fatalCount, "fatal must have called the fatal func exactly once")
		assert.Equal(t, fakeErr, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
	})
}
