package sortFuncs

var SortFuncs = map[string]func([]int) []int{
	"merge":     MergeSort,
	"insertion": InsertionSort,
	"bubble":    BubbleSort,
}
