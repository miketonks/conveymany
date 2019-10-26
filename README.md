# conveymany

TestTables for go-convey


```
func TestIsPrime(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("When IsPrime is called", t, func() {

		ConveyMany("When the number is a prime number, the result should be true",

			TestTable(2, 3, 5, 7, 11, 13, 17, 19),

			func(n int) {
				So(IsPrime(n), ShouldBeTrue)
			})

		ConveyMany("When the number is not a prime number, the result should be false",

			TestTable(1, 4, 6, 8, 9, 10, 12),

			func(n int) {
				So(IsPrime(n), ShouldBeFalse)
			})
	})
}

```
