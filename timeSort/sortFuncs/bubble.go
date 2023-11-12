package sortFuncs

func BubbleSort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if (j + 1) == len(slice) {
				break
			}
			if slice[j] > slice[j+1] {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}
	}
	return slice
}
