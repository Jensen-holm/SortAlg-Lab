package result

import "sync"

type Results struct {
	Data []SortFuncResult
	mu   sync.Mutex
}

func (r *Results) Append(result SortFuncResult) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Data = append(r.Data, result)
}

func NewResults() *Results {
	return &Results{
		mu:   sync.Mutex{},
		Data: []SortFuncResult{},
	}
}
