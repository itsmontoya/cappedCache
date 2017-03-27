package cappedCache

import (
	"github.com/itsmontoya/linkedlist"
)

// New will return a new instance of CappedCache
func New(n int) *CappedCache {
	var c CappedCache
	c.l = linkedlist.New(int32(n), linkedlist.ActionPop)
	return &c
}

// CappedCache is a cache with a capped limit of items
type CappedCache struct {
	l *linkedlist.LinkedList
}

// Insert will add an item to the cache
func (c *CappedCache) Insert(val interface{}) {
	c.l.Append(val)
}

// List will return the current list
func (c *CappedCache) List() (list []interface{}) {
	return c.l.Filter(filterFn)
}

func filterFn(_ interface{}) bool {
	return true
}
