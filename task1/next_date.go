package task1

func NextDate(day, month int) (int, int) {
	odd := []int{1, 3, 5, 7, 8, 10, 12}
	even := []int{4, 6, 9, 11}
	if ((contains(odd, month) && day == 31) || (contains(even, month) && day == 30)) && month != 12 || month == 2 && day == 28 {
		return 1, month + 1
	}
	if month == 12 && day == 31 {
		return 1, 1
	}
	// fmt.Println(day, month)
	return day + 1, month
}

func contains(arr []int, item int) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}
