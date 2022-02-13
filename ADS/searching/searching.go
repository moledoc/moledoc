package searching

// Linear is a function to find the index of an integer in a list
// Time complexity is
//   * worst case O(n)
//   * average case O(n)
//   * best case O(1)
func Linear(A []int, x int) int {
	for i, elem := range A {
		if elem == x {
			return i
		}
	}
	return -1
}

// Binary is a function to find the index of an integer in a sorted list.
// Time complexity is
//   * worst case O(log(n))
//   * average case O(log(n))
//   * best case O(1)
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
