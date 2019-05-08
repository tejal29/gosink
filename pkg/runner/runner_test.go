package runner

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/tejal29/gosink/pkg/worker"
)

func TestRun(t *testing.T) {
	tests := []struct {
		description string
		work        []int
		resultSeen  string
		result      []worker.Result
		isParallel  bool
	}{
		{
			description: "insequence shd see results in sequence",
			work:        []int{1, 3, 2},
			resultSeen:  "1,3,2",
			result:      []worker.Result{1, 3, 2},
			isParallel:  false,
		},
		{
			description: "in parallel shd see results as they are out",
			work:        []int{3, 1, 4},
			resultSeen:  "1,3,4",
			result:      []worker.Result{3, 1, 4},
			isParallel:  true,
		},
		{
			description: "in parallel",
			work:        []int{2, 2, 1, 1},
			resultSeen:  "1,1,2,2",
			result:      []worker.Result{2, 2, 1, 1},
			isParallel:  true,
		},
	}
	for _, test := range tests {
		out := ioutil.Discard
		t.Run(test.description, func(t *testing.T) {
			stream, results := Run(test.isParallel, test.work, out)
			// for testing
			worker.SleepUnit = time.Millisecond
			if stream != test.resultSeen {
				t.Errorf("expected %s, Got %s", test.resultSeen, stream)
			}
			if cmp.Diff(results, test.result) != "" {
				t.Errorf("diff %s", cmp.Diff(results, test.result))
			}
		})
	}
}
