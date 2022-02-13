package sorting

// Insertion is a function that implements insertion-sorting algorithm for integer list.
// Output is increasingly sorted list.
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

// Insertion2 is a funtion that implements insertion-sorting algorithm for integer list.
// Difference with the regular one, is that this one sorts from right to left.
func Insertion2(A []int) []int {
	for j := len(A) - 2; j >= 0; j-- {
		key := A[j]
		i := j + 1
		for i < len(A)-1 && A[i] < key {
			A[i-1] = A[i]
			i++
		}
		A[i-1] = key
	}
	return A
}

// InsertionDesc is a function that implements insertion-sorting algorithm for integer list.
// Output is decreasingly sorted list.
func InsertionDesc(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] < key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
	return A
}
