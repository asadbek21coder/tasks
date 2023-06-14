package task1

func SumArray(a, b float64, n int) (res []float64) {
	res = []float64{0, 0}
	res[0] = a
	res[1] = b
	for i := 2; i < n; i++ {
		var sum float64
		for j := 0; j < i; j++ {
			sum += res[j]
		}
		res = append(res, sum)
	}
	return res
}
