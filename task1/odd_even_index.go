package task1

func OddEvenIndex(arr []int) (odd []int, even []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			even = append(even, arr[i])
		} else {
			odd = append(odd, arr[i])
		}
	}
	return odd, even
}
