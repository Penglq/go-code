package main

import (
	"fmt"

	"go-code/myProtobuf/my"

	"github.com/golang/protobuf/proto"
)

func main() {

	test := &my.TestInt{
		S: 1,
	}
	t, _ := proto.Marshal(test)
	fmt.Printf("int\t%#v\n", t)

	//testSInt := &my.TestSInt{
	//	S: -1,
	//}
	//tSInt, _ := proto.Marshal(testSInt)
	//fmt.Printf("sint负数\t%#v\n", tSInt)
	//
	testString := &my.TestString{
		S: "a",
	}
	tString, _ := proto.Marshal(testString)
	fmt.Printf("字符串\t%#v\n", tString)
	//
	//testFixed32 := &my.TestFixed32{
	//	S: 1,
	//}
	//testF, _ := proto.Marshal(testFixed32)
	//fmt.Printf("fixed32 \t%#v\n", testF)
	//
	testMore := &my.Test{
		Label:         "a",
		Type:          1,
		Reps:          []int64{1, 2},
		RequiredField: "ab",
	}
	tests, _ := proto.Marshal(testMore)
	fmt.Printf("test \t%#v\n", tests)
	//
	//fmt.Println("字节数(byte)", proto.Size(test))
}
