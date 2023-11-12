package sortFuncs

var SortFuncs = map[string]func([]int) []int{
	"merge sort":     MergeSort,
	"insertion sort": InsertionSort,
	"bubble sort":    BubbleSort,
}
