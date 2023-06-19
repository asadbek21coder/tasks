package task1

import "math"

func DifferenceDiskSphere(r1, r2 float64) (diff float64) {
	const pi float64 = 3.1415
	area1 := pi * r1 * r1
	area2 := pi * r2 * r2
	diff = math.Abs(area2 - area1)
	return diff
}
