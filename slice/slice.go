package main

import "log"

func main() {
	// 长度1，容量3
	s := make([]int, 1, 3)
	s[0] = 1
	log.Print("s:", cap(s))

	// 声明之后append长度会增加
	s1 := make([]int, 1)
	s1 = append(s1, 1)
	log.Print("s1:", s1)

	// 创建一个整型切片，元素并没有复制，而是共用
	// 如果更改一个的值，另一个也会变
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	log.Print("slice[1]", &slice[1], &newSlice[0])
	newSlice[0] = 1
	log.Print("值是否改变了", slice[1], newSlice[0])
	newSlice = append(newSlice, 333)
	log.Print("长度和容量", slice, newSlice, len(newSlice), cap(newSlice))

	// 切片append之后，原来的值被释放了吗
	reSlice := []int{1}
	var p *int = &reSlice[0]
	log.Print("切片append之前", &reSlice[0])

	reSlice = append(reSlice, 2)
	log.Print("切片append之后", &reSlice[0], *p)

}
