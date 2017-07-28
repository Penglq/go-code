package main

import (
	"fmt"
	"sync"
	"time"
)

var _self *Singleton

type Singleton struct {
	Name string
	sync.Mutex
}

func NewInstance(s string) *Singleton {
	// 为了出现double instance的并发情况
	time.Sleep(4 * time.Second)
	_self.Name = s
	return _self
}

func Instance(s string) *Singleton {
	// double check
	if _self.Name == "" {
		_self.Lock()
		defer _self.Unlock()
		if _self.Name == "" {
			return NewInstance(s)
		}
	}
	return _self
}

func main() {
	_self = new(Singleton)
	go Instance("a")
	go Instance("b")

	time.Sleep(time.Second * 8)
	fmt.Println(_self.Name)
}
