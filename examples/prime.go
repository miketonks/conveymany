package examples

import "math/big"

// IsPrime checks if a number is a prime number
func IsPrime(i int) bool {
	return big.NewInt(int64(i)).ProbablyPrime(0)
}
