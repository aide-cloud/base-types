package types

import "sync"

type Map struct {
	datasource map[string]interface{}
	length     int
	lock       *sync.RWMutex
}

func (m *Map) Get(key string) (interface{}, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	v, ok := m.datasource[key]
	return v, ok
}

// Set sets the value for the given key.
func (m *Map) Set(key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.datasource[key] = value
	m.length++
}

// Delete deletes the value for the given key.
func (m *Map) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.datasource, key)
	m.length--
}

// Has returns true if the given key is set.
func (m *Map) Has(key string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, ok := m.datasource[key]
	return ok
}

// Keys returns the keys of the map.
func (m *Map) Keys() []string {
	m.lock.RLock()
	defer m.lock.RUnlock()
	keys := make([]string, len(m.datasource))
	i := 0
	for k := range m.datasource {
		keys[i] = k
		i++
	}
	return keys
}

// Values returns the values of the map.
func (m *Map) Values() []interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	values := make([]interface{}, len(m.datasource))
	i := 0
	for _, v := range m.datasource {
		values[i] = v
		i++
	}
	return values
}

// Len returns the number of elements in the map.
func (m *Map) Len() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.length
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.datasource = make(map[string]interface{})
}

// Copy returns a copy of the map.
func (m *Map) Copy() *Map {
	c := NewMap()
	if m.Len() > 0 {
		for k, v := range m.datasource {
			c.Set(k, v)
		}
		c.length = m.length
	}
	return c
}

// NewMap returns a new map.
func NewMap() *Map {
	return &Map{
		datasource: make(map[string]interface{}),
		length:     0,
		lock:       &sync.RWMutex{},
	}
}

// Merge merges the given map into the current one.
func (m *Map) Merge(other *Map, overwrite ...bool) {
	if other == nil {
		return
	}

	if len(overwrite) == 0 {
		overwrite = []bool{false}
	}

	m.lock.Lock()
	defer m.lock.Unlock()
	for k, v := range other.datasource {
		if !m.Has(k) {
			m.Set(k, v)
			m.length++
			continue
		}
		if overwrite[0] {
			m.Set(k, v)
		}
	}
}

// Datasource shows the map.
func (m *Map) Datasource() *map[string]interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return &m.datasource
}
