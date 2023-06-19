package task1

import "fmt"

func OddIncEvenDec(array []int) (odd, even []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		min, index := min(array)
		if min%2 == 0 {
			even = append(even, min)
		} else {
			odd = append(odd, min)
		}
		fmt.Println(array[:index], array[index+1:])
		array = append(array[:index], array[index+1:]...)
	}
	even = reverseArray(even)

	return odd, even
}

func min(array []int) (min int, index int) {
	min = array[0]
	for i := 1; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
			index = i
		}
	}
	return min, index
}

func reverseArray(array []int) (res []int) {
	for i := len(array) - 1; i >= 0; i-- {
		res = append(res, array[i])
	}
	return res
}
