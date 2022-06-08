package workqueue

import (
	"sync"
	"time"

	"k8s.io/utils/clock"
)

// Queue for abstract work queue
type Queue interface {
	Add(item interface{})
	Len() int
	Get() (item interface{}, shutdown bool)
	Done(item interface{})
	Shutdown()
	ShuttingDown() bool
}

type t interface{}
type empty struct{}
type set map[t]empty

func (s set) add(item t) {
	s[item] = empty{}
}

func (s set) del(item t) {
	delete(s, item)
}

func (s set) has(item t) bool {
	_, ok := s[item]
	return ok
}

// FifoQueue implement abstract quene
type FifoQueue struct {
	queue        []t
	dirty        set
	processing   set
	cond         *sync.Cond
	clock        clock.Clock
	shuttingDown bool
	updatePeriod time.Duration
}

// New create a new work queue with default name
func New() *FifoQueue { return NewNamedQueue("default") }

// NewNamedQueue create a new work queue with name
func NewNamedQueue(name string) *FifoQueue {
	c := clock.RealClock{}
	return newQueue(c, time.Millisecond*500)
}

func newQueue(c clock.Clock, updatePeriod time.Duration) *FifoQueue {
	return &FifoQueue{
		clock:        c,
		dirty:        set{},
		processing:   set{},
		updatePeriod: updatePeriod,
		cond:         sync.NewCond(&sync.Mutex{}),
	}
}

// Add make item to be processing
func (f *FifoQueue) Add(item interface{}) {
	f.cond.L.Lock()
	defer f.cond.L.Unlock()
	if f.shuttingDown {
		return
	}

	f.dirty.add(item)
	if f.processing.has(item) {
		return
	}

	f.queue = append(f.queue, item)
	f.cond.Signal()
}

// Len return current quee length
func (f *FifoQueue) Len() int {
	f.cond.L.Lock()
	defer f.cond.L.Unlock()
	return len(f.queue)
}

// Get return an item to be processed
func (f *FifoQueue) Get() (item interface{}, shutdown bool) {
	f.cond.L.Lock()
	defer f.cond.L.Unlock()
	for len(f.queue) == 0 && !f.shuttingDown {
		f.cond.Wait()
	}

	if len(f.queue) == 0 {
		return nil, true
	}

	item = f.queue[0]
	f.queue = f.queue[1:]
	f.processing.add(item)
	f.dirty.del(item)
	return item, false
}

// Done make processing item to be done
func (f *FifoQueue) Done(item interface{}) {
	f.cond.L.Lock()
	defer f.cond.L.Unlock()

	f.processing.del(item)
	if f.dirty.has(item) {
		f.queue = append(f.queue, item)
		f.cond.Signal()
	}
}

// Shutdown close the queue and exit all processing items
func (f *FifoQueue) Shutdown() {
	f.cond.L.Lock()
	defer f.cond.L.Unlock()
	f.shuttingDown = true
	f.cond.Broadcast()
}

// ShuttingDown check current shutting down state
func (f *FifoQueue) ShuttingDown() bool {
	f.cond.L.Lock()
	defer f.cond.L.Unlock()
	return f.shuttingDown
}
