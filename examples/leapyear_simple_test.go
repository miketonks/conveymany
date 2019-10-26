package examples

import (
	"testing"
)

func Test_IsLeapYearBasic(t *testing.T) {

	tests := []int{4, 8, 12, 16, 1996, 2000, 2004, 2008}

	for _, year := range tests {
		if IsLeapYear(year) != true {
			t.Error("IsLeapYear() == true as expected.")
		}
	}
}

func Test_IsLeapYearNegativeCases(t *testing.T) {

	tests := []int{1, 2, 3, 5, 6, 7, 9, 1999, 2001, 2002, 2003, 2005}

	for _, year := range tests {
		if IsLeapYear(year) != false {
			t.Error("IsLeapYear() == false as expected.")
		}
	}
}

func Test_IsLeapYear100Rule(t *testing.T) {

	tests := []int{100, 200, 300, 500, 600, 700, 900, 1900, 2100, 2200, 2300, 2500}

	for _, year := range tests {
		if IsLeapYear(year) != false {
			t.Error("IsLeapYear() == false as expected.")
		}
	}
}

func Test_IsLeapYear400Rule(t *testing.T) {

	tests := []int{400, 800, 1200, 2000, 2400}

	for _, year := range tests {
		if IsLeapYear(year) != true {
			t.Error("IsLeapYear() == true as expected.")
		}
	}
}
