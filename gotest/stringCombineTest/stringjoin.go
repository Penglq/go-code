package stringCombineTest

import (
	"bytes"
	"strings"
)

func StringsJoin(strs []string) string {
	return strings.Join(strs, "")
}

func StringsString(strs []string) string {
	var str string
	s := len(strs)
	for i := 0; i < s; i++ {
		str += strs[i]
	}
	return str
}

func StringBuffer(strs []string) string {
	var buf bytes.Buffer
	s := len(strs)
	for i := 0; i < s; i++ {
		buf.WriteString(strs[i])
	}
	return buf.String()
}

func StringByte(strs []string) string {
	var b2 []byte
	s := len(strs)
	for i := 0; i < s; i++ {
		b2 = append(b2,[]byte(strs[i])...)
	}
	return string(b2)
}
