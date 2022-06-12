package workqueue

import (
	"sync"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	workq := New()

	producers := 100
	prodWg := sync.WaitGroup{}
	for i := 0; i < producers; i++ {
		prodWg.Add(1)
		go func(i int) {
			defer prodWg.Done()
			workq.Add(i)

			// * Test Duplicate Add
			if i == producers {
				workq.Add(i)
			}
		}(i)
	}

	consumers := 15
	consWg := sync.WaitGroup{}
	for i := 0; i < consumers; i++ {
		consWg.Add(1)
		go func(i int) {
			defer consWg.Done()
			for {
				item, exit := workq.Get()
				if exit {
					return
				}

				if item == "added item after shutdown!" {
					t.Error("got an item after queue shutdown")
				}

				t.Logf("Worker-%03v: begin processing %v", i, item)
				time.Sleep(time.Millisecond * 1000)
				t.Logf("Worker-%03v: finish processed %v", i, item)
				workq.Done(item)
				t.Logf("current wrokqueue length: %v", workq.Len())
			}
		}(i)
	}

	t.Logf("workqueue current shutdown: %v", workq.ShuttingDown())
	prodWg.Wait()
	workq.Shutdown()
	workq.Add("added item after shutdown!")
	consWg.Wait()
	t.Logf("workqueue current shutdown: %v", workq.ShuttingDown())
}
