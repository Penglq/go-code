// Golang的单例模式该怎么写？随手写一个，不错，立马写出来了。但这个代码有什么问题呢？
// 多个协程同时执行这段代码就会出现问题：instance可能会被赋值多次，这段代码是线程不安全的代码。
// 那么如何保证在多线程下只执行一次呢？条件反射：加锁。。。加锁是可以解决问题。
// 但不是最优的方案，因为如果有1W并发，每一个线程都竞争锁，同一时刻只有一个线程能拿到锁，其他的全部阻塞等待。
// 让原本想并发得飞起来变成了一切认怂串行化。通过check-lock-check方式可以减少竞争。还有其他方式，利用sync/atomic和sync/once 这里只给出代码

package singleton

import (
	"sync"
	"sync/atomic"
)

var instance *singleton
var initialized *uint32

type singleton struct{}

// 并发不安全
func NewSingletonOne() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// 加锁，多个协程程序变成串行，大家竞争一个锁
func NewSingletonTwo() *singleton {
	mux := sync.Mutex{}
	mux.Lock() // lock
	defer mux.Unlock()
	if instance == nil { // check
		instance = &singleton{}
	}
	return instance
}

// check-lock-check方式加锁，减少锁竞争
func NewSingletonThree() *singleton {
	if instance == nil { // check
		mux := sync.Mutex{}
		mux.Lock() // lock
		defer mux.Unlock()
		if instance == nil { // check
			instance = &singleton{}
		}
	}
	return instance
}

// sync/atomic使用原子方式
func NewSingletonFour() *singleton {
	if atomic.LoadUint32(initialized) == 1 {
		return &singleton{}
	}
	mux := sync.Mutex{}
	mux.Lock() // lock
	defer mux.Unlock()
	if *initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(initialized, 1)
	}
	return instance
}

// sync/once
func NewSingletonFive() *singleton {
	sync.Once{}.Do(func() {
		instance = &singleton{}
	})
	return instance
}