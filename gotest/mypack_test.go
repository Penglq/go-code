package gotest

import "testing"

// 输入的参数列表
var fibTests = []struct {
	a        float64 // input
	b        float64 // input
	expected float64 // expected result
}{
	{4, 2, 2},
	{6, 3, 2},
	{100, 10, 10},
	{8, 2, 4},
	{15, 5, 3},
}

func Test_Division_1(t *testing.T) {
	for _, fib := range fibTests {
		if i, e := Division(fib.a, fib.b); i != fib.expected || e != nil { //try a unit test on function
			t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
		} else {
			t.Log("第一个测试通过了") //记录一些你期望记录的信息
		}
	}
}

func Test_Division_2(t *testing.T) {
	if _, e := Division(6, 0); e == nil { //try a unit test on function
		t.Error("Division did not work as expected.") // 如果不是如预期的那么就报错
	} else {
		t.Log("one test passed.", e) //记录一些你期望记录的信息
	}
}
