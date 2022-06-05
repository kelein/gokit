package cmap

import (
	"sync"
)

// IntMap for integer
type IntMap struct {
	m sync.Map
}

// Store save key value pairs
func (i *IntMap) Store(key int, value string) {
	i.m.Store(key, value)
}

// Load get key value pair
func (i *IntMap) Load(key int) (string, bool) {
	v, ok := i.m.Load(key)
	if v != "" {
		return v.(string), ok
	}
	return "", false
}

// LoadOrStore get or save key value pair
func (i *IntMap) LoadOrStore(key int, value string) (string, bool) {
	v, ok := i.m.LoadOrStore(key, value)
	return v.(string), ok
}

// Range iterate key value pairs
func (i *IntMap) Range(f func(key int, value string) bool) {
	fn := func(key, value interface{}) bool {
		return f(key.(int), value.(string))
	}
	i.m.Range(fn)
}

// Delete destory key value pair
func (i *IntMap) Delete(key int) {
	i.m.Delete(key)
}
