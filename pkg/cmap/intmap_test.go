package cmap

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestIntMap(t *testing.T) {
	type args struct{ key int }
	tests := []struct{ name string }{
		{"A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			im := &IntMap{}

			// * Test Store
			wg := sync.WaitGroup{}
			for i := 0; i < 10000; i++ {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					im.Store(i, time.Now().String())
				}(i)
			}

			// * Test Range
			wg.Add(1)
			go func() {
				defer wg.Done()
				fn := func(key int, value string) bool {
					log.Printf("<%d: %s>", key, value)
					return true
				}
				im.Range(fn)
			}()

			wg.Wait()
		})
	}
}
