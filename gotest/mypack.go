package gotest

import (
	"errors"
	"time"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	time.Sleep(1 * time.Second)
	return a / b, nil
}
