package insertion

// Insertion is a function that implements insertion-sorting algorithm for integer list.
// Time complexity is
//   * worst case O(n^2)
//   * average case O(n^2)
//   * best case O(n)
func Insertion(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
	return A
}
