package binary

func Binary(A []int, x int) int {
	if len(A) == 0 {
		return -1
	}
	l := 0
	r := len(A) - 1
	for l < r {
		m := (l + r) / 2
		if A[m] == x {
			return m
		}
		if A[m] < x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if A[l] == x {
		return l
	}
	return -1
}
