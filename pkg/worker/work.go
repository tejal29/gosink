package worker

import (
	"sync"
	"time"
)

// Result is work result
type Result int

var SleepUnit = time.Second

// InSequence does work in sequence and waits for the previous work
// item to complete. it should write results on channel in order of
// input work array.
func InSequence(ch chan Result, work []int) []Result {
	results := make([]Result, len(work))
	go buildInSeq(ch, work, results)
	return results
}

// InParallel does work in parallel and should write results on channels
//as work item complete
func InParallel(ch chan Result, work []int) []Result {
	results := make([]Result, len(work))
	var mutex = &sync.Mutex{}
	for i, num := range work {
		go func(w int, res *Result, ch chan Result) {
			*res = executeWork(w)
			mutex.Lock()
			ch <- *res
			mutex.Unlock()
		}(num, &results[i], ch)
	}
	return results
}

// build sleep for time mentioned in num and returns the same number
func executeWork(num int) Result {
	time.Sleep(time.Duration(num) * SleepUnit)
	return Result(num)
}

func buildInSeq(ch chan Result, work []int, results []Result) {
	for i, num := range work {
		res := executeWork(num)
		ch <- res
		results[i] = res
	}
}
