package main

import (
	"container/list"
	"fmt"
)

func main() {
	//初始化一个list
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushFront(5)

	fmt.Println("Before Removing...")

	var n *list.Element
	//遍历list，删除元素
	for e := l.Front(); e != nil; e = n {
		fmt.Println("removing", e.Value)
		n = e.Next()
		l.Remove(e)
	}
	fmt.Println("After Removing...")
	//遍历删除完元素后的list
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
