package examples

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Convey_IsLeapYear(t *testing.T) {

	Convey("When IsLeapYear ia called", t, func() {

		Convey("When the year is divisible by 4, but not by 100", func() {

			tests := []int{4, 8, 12, 16, 1996, 2000, 2004, 2008}

			for _, year := range tests {
				So(IsLeapYear(year), ShouldBeTrue)
			}
		})

		Convey("When the year is not divisible by 4", func() {

			tests := []int{1, 2, 3, 5, 6, 7, 9, 1999, 2001, 2002, 2003, 2005}

			for _, year := range tests {
				So(IsLeapYear(year), ShouldBeFalse)
			}
		})

		Convey("When the year is divisible by 100, but not 400", func() {

			tests := []int{100, 200, 300, 500, 600, 700, 900, 1900, 2100, 2200, 2300, 2500}

			for _, year := range tests {
				So(IsLeapYear(year), ShouldBeFalse)
			}
		})

		Convey("When the year is divisible by 400", func() {

			tests := []int{400, 800, 1200, 2000, 2400}

			for _, year := range tests {
				So(IsLeapYear(year), ShouldBeTrue)
			}
		})
	})
}
