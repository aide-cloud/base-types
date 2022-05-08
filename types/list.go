package types

import (
	"sort"
	"sync"
)

// NoIndex 不存在的index
const NoIndex = -1

type List struct {
	items  []interface{}
	length int
	lock   *sync.RWMutex
}

// NewList creates a new list with the given items.
func NewList(items ...interface{}) *List {
	return &List{
		items:  items,
		length: len(items),
		lock:   &sync.RWMutex{},
	}
}

// Len returns the length of the list.
func (l *List) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.length
}

// Values returns the items in the list.
func (l *List) Values() []interface{} {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.items
}

// Append adds the given items to the list.
func (l *List) Append(items ...interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.items = append(l.items, items...)
	l.length += len(items)
}

// RemoveByValue removes the given items from the list.
func (l *List) RemoveByValue(items ...interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	for _, item := range items {
		for i, item2 := range l.items {
			if item == item2 {
				l.items = append(l.items[:i], l.items[i+1:]...)
				l.length--
			}
		}
	}
}

// Sort sorts the list.
func (l *List) Sort(less func(i, j int) bool) {
	l.lock.Lock()
	defer l.lock.Unlock()

	sort.Slice(l.items, less)
}

// Clear clears the list.
func (l *List) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.items = []interface{}{}
	l.length = 0
}

// Contains returns true if the given item is in the list.
func (l *List) Contains(item interface{}) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	for _, item2 := range l.items {
		if item == item2 {
			return true
		}
	}
	return false
}

// Index returns the index of the given item.
func (l *List) Index(item interface{}) int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	for i, item2 := range l.items {
		if item == item2 {
			return i
		}
	}
	return NoIndex
}

// Indexes returns the indexes of the given items.
func (l *List) Indexes(items ...interface{}) []int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	var indexes []int
	for _, item := range items {
		for i, item2 := range l.items {
			if item == item2 {
				indexes = append(indexes, i)
			}
		}
	}
	return indexes
}

// ContainsAll returns true if all the given items are in the list.
func (l *List) ContainsAll(items ...interface{}) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	for _, item := range items {
		if !l.Contains(item) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if any of the given items are in the list.
func (l *List) ContainsAny(items ...interface{}) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	count := 0
	for _, item := range items {
		if l.Contains(item) {
			count++
		}
	}
	return count == len(items)
}

// ContainsNone returns true if none of the given items are in the list.
func (l *List) ContainsNone(items ...interface{}) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	for _, item := range items {
		if l.Contains(item) {
			return false
		}
	}
	return true
}

// Unique returns a list with unique items.
func (l *List) Unique() {
	l.lock.Lock()
	defer l.lock.Unlock()
	var items = NewList()
	count := 0
	for _, item := range l.items {
		if !items.Contains(item) {
			items.Append(item)
			count++
		}
	}
	l.items = items.Values()
	l.length = count
}

// Get returns the item at the given index.
func (l *List) Get(index int) interface{} {
	l.lock.RLock()
	defer l.lock.RUnlock()
	if index < 0 || index >= l.length {
		return nil
	}
	return l.items[index]
}

// Set sets the item at the given index.
func (l *List) Set(index int, item interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if index < 0 || index >= l.length {
		return
	}
	l.items[index] = item
}

// First returns the first item.
func (l *List) First() interface{} {
	l.lock.RLock()
	defer l.lock.RUnlock()
	if l.length == 0 {
		return nil
	}
	return l.items[0]
}

// Last returns the last item.
func (l *List) Last() interface{} {
	l.lock.RLock()
	defer l.lock.RUnlock()
	if l.length == 0 {
		return nil
	}
	return l.items[l.length-1]
}

// RemoveByIndex removes the item at the given index.
func (l *List) RemoveByIndex(index int) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if index < 0 || index >= l.length {
		return
	}
	l.items = append(l.items[:index], l.items[index+1:]...)
	l.length--
}

// Reverse reverses the list.
func (l *List) Reverse() {
	l.lock.Lock()
	defer l.lock.Unlock()
	for i, j := 0, l.length-1; i < j; i, j = i+1, j-1 {
		l.items[i], l.items[j] = l.items[j], l.items[i]
	}
}

// Find returns the first item that satisfies the given condition.
func (l *List) Find(item interface{}) int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	for i, item2 := range l.items {
		if item == item2 {
			return i
		}
	}
	return NoIndex
}
