package task1

func EvenIndexDesc(arr []int) (res []int, count int) {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i]%2 == 0 {
			res = append(res, i)
			count++
		}
		// fmt.Println(i, arr[i])
	}
	return res, count
}
