package task1

import "math"

func IsPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	// if n<=1 {
	// 	return false
	// }
	// return true

	return n > 1
}
