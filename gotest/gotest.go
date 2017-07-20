package gotest

import (
	"errors"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	for i := 0; i < 100005; i++ {
		_ = 123123
	}
	return a / b, nil
}
