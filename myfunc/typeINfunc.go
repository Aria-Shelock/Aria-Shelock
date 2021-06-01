package myfunc

import "errors"

type a func(int, int) int

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Calc(x, y int, op a) int {
	return op(x, y)
}

func Do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return Add, nil
	case "-":
		return Sub, nil
	default:
		err := errors.New("Error")
		return nil, err
	}
}
