package task1

func Kabisa(year int) (res bool) {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		res = true
	}
	return res
}
