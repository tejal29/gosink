package runner

import (
	"fmt"
	"io"
	"strings"

	"github.com/tejal29/gosink/pkg/worker"
)

// Run executes the work in w and returns the order
// the work was done as well as results.
func Run(inParallel bool, w []int, out io.Writer) (string, []worker.Result) {

	in := make(chan worker.Result, 1)
	defer close(in)
	expectedWork := len(w)

	// Exectue work in sequence or parallel
	results := doWork(inParallel, w, in)
	var stream = make([]string, expectedWork)

	// barrier function which waits for all the work to be done
	for i := 0; i < expectedWork; i++ {
		res := <-in
		stream[i] = fmt.Sprintf("%d", int(res))
		out.Write([]byte(stream[i]))
	}
	return strings.Join(stream, ","), results
}

func doWork(parallel bool, w []int, in chan worker.Result) []worker.Result {
	if parallel {
		return worker.InParallel(in, w)
	} else {
		return worker.InSequence(in, w)
	}
}
