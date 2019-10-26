package examples

import (
	"testing"

	. "github.com/miketonks/conveymany"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSum(t *testing.T) {

	// Only pass t into top-level Convey calls
	ConveyMany("When two integers are added together", t,
		TestTable(
			TestValues(1, 1, 2),
			TestValues(1, 2, 3),
			TestValues(5, 6, 11),
			TestValues(1, 3, 4),
			// TestValues(1, 3, 9),
		),
		func(a, b, c int) {
			ans := Sum(a, b)
			So(ans, ShouldEqual, c)
		})
}
