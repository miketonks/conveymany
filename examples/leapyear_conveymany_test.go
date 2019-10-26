package examples

import (
	"testing"

	. "github.com/miketonks/conveymany"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ConveyMany_IsLeapYear(t *testing.T) {

	Convey("When IsLeapYear ia called", t, func() {

		ConveyMany("When the year is divisible by 4, but not by 100",

			TestTable(4, 8, 12, 16, 1996, 2000, 2004, 2008),

			func(year int) {
				So(IsLeapYear(year), ShouldBeTrue)
			})

		ConveyMany("When the year is not divisible by 4",

			TestTable(1, 2, 3, 5, 6, 7, 9, 1999, 2001, 2002, 2003, 2005),

			func(year int) {
				So(IsLeapYear(year), ShouldBeFalse)
			})

		ConveyMany("When the year is divisible by 100, but not 400",

			TestTable(100, 200, 300, 500, 600, 700, 900, 1900, 2100, 2200, 2300, 2500),

			func(year int) {
				So(IsLeapYear(year), ShouldBeFalse)
			})

		ConveyMany("When the year is divisible by 400",

			TestTable(400, 800, 1200, 2000, 2400),

			func(year int) {
				So(IsLeapYear(year), ShouldBeTrue)
			})
	})
}
