package task1

func Descriminant(a, b, c float64) int {
	if b*b-4*a*c > 0 {
		return 2
	} else if b*b-4*a*c < 0 {
		return 0
	}
	return 1
}
