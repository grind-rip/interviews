package main

// Complete the 'getMinMoves' function below.
//
// The function is expected to return an INTEGER.
// The function accepts INTEGER_ARRAY plates as a parameter.

// An inversion occurs when a larger element appears before a smaller element
// in the array. Each inversion must be resolved by at least one adjacent swap
// to achieve sorted order. Conversely, each adjacent swap can resolve at most
// one inversion. Therefore, the minimum number of swaps needed equals the
// number of inversions.
//
// Merge sort is an efficient sorting algorithm that works by repeatedly
// dividing the array in half, sorting each half, and then merging the sorted
// halves. We can modify the merge step to count inversions: when we merge two
// sorted halves, if an element from the right half is chosen, it means all
// remaining elements in the left half are greater than it, creating
// inversions.
//
// Runtime Complexity: O(n log n)
// Space Complexity: O(n) for the temporary arrays used in merge sort

// getMinMovesNaive returns the minimum number of moves required put the
// elements of the array in ascending order using bubble sort.
func getMinMovesNaive(plates []int32) int32 {
	// The critical piece of information is,
	//
	//   >In one move, you can swap the order of adjacent plates.
	//
	// Given this constraint, this requires an implementation of bubble sort,
	// which steps through the list, compares adjacent elements and swaps them if
	// they are in the wrong order.
	moves := 0
	for i := 0; i < len(plates)-1; i++ {
		for j := 0; j < len(plates)-i-1; j++ {
			if plates[j] > plates[j+1] {
				plates[j], plates[j+1] = plates[j+1], plates[j]
				moves++
			}
		}
	}
	return int32(moves)
}

// getMinMovesOptimal returns the minimum number of moves required put the
// elements of the array in ascending order using merge sort.
func getMinMovesOptimal(plates []int32) int32 {
	// Use a modified merge sort such that every time an unsorted pair is found,
	// increment the inversion count, then return the inversion count at the end.
	//
	// NOTE: It is cumbersome to implement this using constant space complexity,
	// that is, O(1). This requires an implementation of an in-place sorting
	// algorithm. Here, we simply ignore the returned, sorted slice.
	_, moves := mergeSort(plates)
	return moves
}

// mergeSort recursively sorts an unsorted array, while counting inversions.
func mergeSort(arr []int32) ([]int32, int32) {
	if len(arr) == 1 {
		return arr, 0
	}
	mid := len(arr) / 2
	l, cl := mergeSort(arr[:mid])
	r, cr := mergeSort(arr[mid:])
	s, cm := merge(l, r)
	return s, (cl + cr + cm)
}

// merge merges two sorted arrays and counts inversions.
//
// At any step in merge(), if arr1[i] is greater than arr2[j], then there are
// (mid – i) inversions.
//
// When we choose an element from the right array (arr2), it means this element
// should come before all remaining elements in the left array (arr1).
// size/2 represents the length of arr1, and i is how many elements we've
// already processed from arr1, so (size/2 - i) gives us the count of remaining
// elements in arr1 that form inversions with the current arr2[j] element.
func merge(arr1 []int32, arr2 []int32) ([]int32, int32) {
	size := len(arr1) + len(arr2)
	sorted := make([]int32, size)
	i, j, k, c := 0, 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			sorted[k] = arr1[i]
			i, k = i+1, k+1
		} else {
			sorted[k] = arr2[j]
			j, k = j+1, k+1
			// Increment inversion count by mid - i, where mid - i ≡ size/2 - i
			c += (size / 2) - i
		}
	}
	for i < len(arr1) {
		sorted[k] = arr1[i]
		i, k = i+1, k+1
	}
	for j < len(arr2) {
		sorted[k] = arr2[j]
		j, k = j+1, k+1
	}
	return sorted, int32(c)
}
