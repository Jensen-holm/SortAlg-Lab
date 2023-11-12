package result

type SortFuncResult struct {
	FuncName string
	Time     float64
	N        int
}

func NewResult(fName string, wt float64, n int) SortFuncResult {
	return SortFuncResult{
		FuncName: fName,
		Time:     wt,
		N:        n,
	}
}
