package task1

func PerfectNumber(n int) (arr []int) {
	for i := 1; i <= n; i++ {
		// fmt.Println(i)
		var sum int
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if sum == i {
			arr = append(arr, i)
		}
	}
	return arr
}
