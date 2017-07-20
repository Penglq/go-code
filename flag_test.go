package main

import (
	"flag"
	"testing"
)

var strFlag = flag.String("s", "", "Description")
var boolFlag = flag.Bool("bool", false, "Description of flag")

func Test_Flag(t *testing.T) {
	flag.Parse()
	println(*strFlag, *boolFlag)
}
func main() {

}
