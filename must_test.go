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

func TestDo3(t *testing.T) {
	t.Run("When does not have error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		fakeValue2 := gofakeit.SentenceSimple()
		fakeValue3 := gofakeit.SentenceSimple()
		value1, value2, value3 := Do3(fakeValue1, fakeValue2, fakeValue3, nil)

		assert.Zero(t, fatalCount, "fatal must have never been called")
		assert.Nil(t, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
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
		fakeValue3 := gofakeit.SentenceSimple()
		fakeErr := errors.New(gofakeit.SentenceSimple())
		value1, value2, value3 := Do3(fakeValue1, fakeValue2, fakeValue3, fakeErr)

		assert.Equal(t, 1, fatalCount, "fatal must have called the fatal func exactly once")
		assert.Equal(t, fakeErr, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
	})
}

func TestDo4(t *testing.T) {
	t.Run("When does not have error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		fakeValue2 := gofakeit.SentenceSimple()
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		value1, value2, value3, value4 := Do4(fakeValue1, fakeValue2, fakeValue3, fakeValue4, nil)

		assert.Zero(t, fatalCount, "fatal must have never been called")
		assert.Nil(t, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
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
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeErr := errors.New(gofakeit.SentenceSimple())
		value1, value2, value3, value4 := Do4(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeErr)

		assert.Equal(t, 1, fatalCount, "fatal must have called the fatal func exactly once")
		assert.Equal(t, fakeErr, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
	})
}

func TestDo5(t *testing.T) {
	t.Run("When does not have error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		fakeValue2 := gofakeit.SentenceSimple()
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		value1, value2, value3, value4, value5 := Do5(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, nil)

		assert.Zero(t, fatalCount, "fatal must have never been called")
		assert.Nil(t, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
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
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		fakeErr := errors.New(gofakeit.SentenceSimple())
		value1, value2, value3, value4, value5 := Do5(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, fakeErr)

		assert.Equal(t, 1, fatalCount, "fatal must have called the fatal func exactly once")
		assert.Equal(t, fakeErr, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
	})
}

func TestDo6(t *testing.T) {
	t.Run("When does not have error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		fakeValue2 := gofakeit.SentenceSimple()
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		fakeValue6 := gofakeit.SentenceSimple()
		value1, value2, value3, value4, value5, value6 := Do6(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, fakeValue6, nil)

		assert.Zero(t, fatalCount, "fatal must have never been called")
		assert.Nil(t, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
		assert.Equal(t, fakeValue6, value6)
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
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		fakeValue6 := gofakeit.SentenceSimple()
		fakeErr := errors.New(gofakeit.SentenceSimple())
		value1, value2, value3, value4, value5, value6 := Do6(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, fakeValue6, fakeErr)

		assert.Equal(t, 1, fatalCount, "fatal must have called the fatal func exactly once")
		assert.Equal(t, fakeErr, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
		assert.Equal(t, fakeValue6, value6)
	})
}

func TestDo7(t *testing.T) {
	t.Run("When does not have error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		fakeValue2 := gofakeit.SentenceSimple()
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		fakeValue6 := gofakeit.SentenceSimple()
		fakeValue7 := gofakeit.SentenceSimple()
		value1, value2, value3, value4, value5, value6, value7 := Do7(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, fakeValue6, fakeValue7, nil)

		assert.Zero(t, fatalCount, "fatal must have never been called")
		assert.Nil(t, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
		assert.Equal(t, fakeValue6, value6)
		assert.Equal(t, fakeValue7, value7)
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
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		fakeValue6 := gofakeit.SentenceSimple()
		fakeValue7 := gofakeit.SentenceSimple()
		fakeErr := errors.New(gofakeit.SentenceSimple())
		value1, value2, value3, value4, value5, value6, value7 := Do7(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, fakeValue6, fakeValue7, fakeErr)

		assert.Equal(t, 1, fatalCount, "fatal must have called the fatal func exactly once")
		assert.Equal(t, fakeErr, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
		assert.Equal(t, fakeValue6, value6)
		assert.Equal(t, fakeValue7, value7)
	})
}

func TestDo8(t *testing.T) {
	t.Run("When does not have error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		fakeValue2 := gofakeit.SentenceSimple()
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		fakeValue6 := gofakeit.SentenceSimple()
		fakeValue7 := gofakeit.SentenceSimple()
		fakeValue8 := gofakeit.SentenceSimple()
		value1, value2, value3, value4, value5, value6, value7, value8 := Do8(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, fakeValue6, fakeValue7, fakeValue8, nil)

		assert.Zero(t, fatalCount, "fatal must have never been called")
		assert.Nil(t, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
		assert.Equal(t, fakeValue6, value6)
		assert.Equal(t, fakeValue7, value7)
		assert.Equal(t, fakeValue8, value8)
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
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		fakeValue6 := gofakeit.SentenceSimple()
		fakeValue7 := gofakeit.SentenceSimple()
		fakeValue8 := gofakeit.SentenceSimple()
		fakeErr := errors.New(gofakeit.SentenceSimple())
		value1, value2, value3, value4, value5, value6, value7, value8 := Do8(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, fakeValue6, fakeValue7, fakeValue8, fakeErr)

		assert.Equal(t, 1, fatalCount, "fatal must have called the fatal func exactly once")
		assert.Equal(t, fakeErr, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
		assert.Equal(t, fakeValue6, value6)
		assert.Equal(t, fakeValue7, value7)
		assert.Equal(t, fakeValue8, value8)
	})
}

func TestDo9(t *testing.T) {
	t.Run("When does not have error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		fakeValue2 := gofakeit.SentenceSimple()
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		fakeValue6 := gofakeit.SentenceSimple()
		fakeValue7 := gofakeit.SentenceSimple()
		fakeValue8 := gofakeit.SentenceSimple()
		fakeValue9 := gofakeit.SentenceSimple()
		value1, value2, value3, value4, value5, value6, value7, value8, value9 := Do9(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, fakeValue6, fakeValue7, fakeValue8, fakeValue9, nil)

		assert.Zero(t, fatalCount, "fatal must have never been called")
		assert.Nil(t, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
		assert.Equal(t, fakeValue6, value6)
		assert.Equal(t, fakeValue7, value7)
		assert.Equal(t, fakeValue8, value8)
		assert.Equal(t, fakeValue9, value9)
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
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		fakeValue6 := gofakeit.SentenceSimple()
		fakeValue7 := gofakeit.SentenceSimple()
		fakeValue8 := gofakeit.SentenceSimple()
		fakeValue9 := gofakeit.SentenceSimple()
		fakeErr := errors.New(gofakeit.SentenceSimple())
		value1, value2, value3, value4, value5, value6, value7, value8, value9 := Do9(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, fakeValue6, fakeValue7, fakeValue8, fakeValue9, fakeErr)

		assert.Equal(t, 1, fatalCount, "fatal must have called the fatal func exactly once")
		assert.Equal(t, fakeErr, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
		assert.Equal(t, fakeValue6, value6)
		assert.Equal(t, fakeValue7, value7)
		assert.Equal(t, fakeValue8, value8)
		assert.Equal(t, fakeValue9, value9)
	})
}

func TestDo10(t *testing.T) {
	t.Run("When does not have error", func(t *testing.T) {
		var fatalCount int
		var fatalErr error
		SetFatalFunc(func(err error) {
			fatalCount++
			fatalErr = err
		})

		fakeValue1 := gofakeit.SentenceSimple()
		fakeValue2 := gofakeit.SentenceSimple()
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		fakeValue6 := gofakeit.SentenceSimple()
		fakeValue7 := gofakeit.SentenceSimple()
		fakeValue8 := gofakeit.SentenceSimple()
		fakeValue9 := gofakeit.SentenceSimple()
		fakeValue10 := gofakeit.SentenceSimple()
		value1, value2, value3, value4, value5, value6, value7, value8, value9, value10 := Do10(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, fakeValue6, fakeValue7, fakeValue8, fakeValue9, fakeValue10, nil)

		assert.Zero(t, fatalCount, "fatal must have never been called")
		assert.Nil(t, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
		assert.Equal(t, fakeValue6, value6)
		assert.Equal(t, fakeValue7, value7)
		assert.Equal(t, fakeValue8, value8)
		assert.Equal(t, fakeValue9, value9)
		assert.Equal(t, fakeValue10, value10)
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
		fakeValue3 := gofakeit.SentenceSimple()
		fakeValue4 := gofakeit.SentenceSimple()
		fakeValue5 := gofakeit.SentenceSimple()
		fakeValue6 := gofakeit.SentenceSimple()
		fakeValue7 := gofakeit.SentenceSimple()
		fakeValue8 := gofakeit.SentenceSimple()
		fakeValue9 := gofakeit.SentenceSimple()
		fakeValue10 := gofakeit.SentenceSimple()
		fakeErr := errors.New(gofakeit.SentenceSimple())
		value1, value2, value3, value4, value5, value6, value7, value8, value9, value10 := Do10(fakeValue1, fakeValue2, fakeValue3, fakeValue4, fakeValue5, fakeValue6, fakeValue7, fakeValue8, fakeValue9, fakeValue10, fakeErr)

		assert.Equal(t, 1, fatalCount, "fatal must have called the fatal func exactly once")
		assert.Equal(t, fakeErr, fatalErr)
		assert.Equal(t, fakeValue1, value1)
		assert.Equal(t, fakeValue2, value2)
		assert.Equal(t, fakeValue3, value3)
		assert.Equal(t, fakeValue4, value4)
		assert.Equal(t, fakeValue5, value5)
		assert.Equal(t, fakeValue6, value6)
		assert.Equal(t, fakeValue7, value7)
		assert.Equal(t, fakeValue8, value8)
		assert.Equal(t, fakeValue9, value9)
		assert.Equal(t, fakeValue10, value10)
	})
}
