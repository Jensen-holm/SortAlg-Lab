package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	r "sort/result"
	s "sort/sortFuncs"
	"sync"
	"time"
)

func main() {
	var (
		nVals   = [6]int{2048, 4096, 8192, 16384, 32768, 65536}
		results = r.NewResults()
		wg      sync.WaitGroup
	)

	for _, n := range nVals {
		wg.Add(len(s.SortFuncs))
		slice := generateSlice(n)

		for funcName, sortFunc := range s.SortFuncs {
			funcName := funcName
			sortFunc := sortFunc
			go func(funcName string, sortFunc func([]int) []int, n int) {
				defer wg.Done()
				// make a copy of the slice
				sCopy := make([]int, n)
				copy(sCopy, slice)

				// perform and time the sortfunc
				st := time.Now()
				_ = sortFunc(sCopy)
				elapsed := time.Since(st).Seconds()

				// store results in result struct
				result := r.NewResult(funcName, elapsed, n)
				results.Append(result)
			}(funcName, sortFunc, n)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	if len(results.Data) != (len(s.SortFuncs) * len(nVals)) {
		err := errors.New("incorrect computation in main function")
		log.Fatalf("%s", err)
	}

	for i := 0; i < len(results.Data); i++ {
		result := results.Data[i]
		fmt.Printf("\nfunc: %s\nwall time: %f\nn: %d\n", result.FuncName, result.Time, result.N)
	}

	err := saveResults(results)
	if err != nil {
		panic(err)
	}

}

func generateSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func saveResults(results *r.Results) error {
	path := "../resultData/results.csv"

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Function", "Wall Time", "N"}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	// Write data
	for _, result := range results.Data {
		row := []string{result.FuncName, fmt.Sprintf("%f", result.Time), fmt.Sprintf("%d", result.N)}
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
