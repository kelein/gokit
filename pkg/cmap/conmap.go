package cmap

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

var (
	errInvalidKeyType = errors.New("invalid key type")
	errInvalidValType = errors.New("invalid value type")
)

// ConMap for concurrency
type ConMap struct {
	store sync.Map
	ktype reflect.Type
	vtype reflect.Type
}

// NewConMap create a ConMap instance
func NewConMap(ktype, vtype reflect.Type) (*ConMap, error) {
	if ktype == nil || vtype == nil {
		return nil, errors.New("type nil error")
	}
	if !ktype.Comparable() {
		return nil, fmt.Errorf("uncomparable key: %v", ktype)
	}
	return &ConMap{ktype: ktype, vtype: vtype}, nil
}

// Store save key value pairs into ConMap
func (c *ConMap) Store(key, value interface{}) error {
	if err := c.validKey(key); err != nil {
		return fmt.Errorf("%v: %v", err, key)
	}
	if err := c.validValue(value); err != nil {
		return fmt.Errorf("%v: %v", err, value)
	}
	c.store.Store(key, value)
	return nil
}

// Load get key value pair from ConMap
func (c *ConMap) Load(key interface{}) (interface{}, bool) {
	if err := c.validKey(key); err != nil {
		return nil, false
	}
	return c.store.Load(key)
}

// LoadOrStore get key value pair from ConMap or store it if not exists
func (c *ConMap) LoadOrStore(key, value interface{}) (interface{}, bool) {
	if err := c.validKey(key); err != nil {
		return nil, false
	}
	if err := c.validValue(value); err != nil {
		return nil, false
	}
	return c.store.LoadOrStore(key, value)
}

// Range visit each key in ConMap
func (c *ConMap) Range(f func(key, value interface{}) bool) {
	c.store.Range(f)
}

// Delete destory key value pair from ConMap
func (c *ConMap) Delete(key interface{}) error {
	if err := c.validKey(key); err != nil {
		return fmt.Errorf("%v: %v", err, key)
	}
	c.store.Delete(key)
	return nil
}

func (c *ConMap) validKey(key interface{}) error {
	if reflect.TypeOf(key) != c.ktype {
		return errInvalidKeyType
	}
	return nil
}

func (c *ConMap) validValue(value interface{}) error {
	if reflect.TypeOf(value) != c.vtype {
		return errInvalidValType
	}
	return nil
}
