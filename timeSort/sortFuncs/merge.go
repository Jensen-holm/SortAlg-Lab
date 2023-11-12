package sortFuncs

func MergeSort(slice []int) []int {
	if len(slice) == 1 {
		return slice
	}

	mid := len(slice) / 2
	left := slice[:mid]
	right := slice[mid:]
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
