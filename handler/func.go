package handler

// Abs absolute value for ints
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
