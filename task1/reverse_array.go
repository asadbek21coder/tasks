package task1

func ReverseArray(arr []int) (res []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}
	return res
}

// func Append(arr []int, item int) []int {
// 	arr = append(arr, item)
// 	return arr
// }
