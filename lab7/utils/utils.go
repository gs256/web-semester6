package utils

func Clamp(value int, min int, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}
