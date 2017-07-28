package stringutil

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some ASCII and UTF8 strings", t, func() {
		strs := []struct {
			in, want string
		}{
			{"Hello, world", "dlrow ,olleH"},
			{"Hello, 世界", "界世 ,olleH"},
			{"", ""},
		}
		Convey("The value should be equal the reversed one", func() {
			for _, c := range strs {
				got := Reverse(c.in)
				So(got, ShouldEqual, c.want)
			}
		})
	})
}
