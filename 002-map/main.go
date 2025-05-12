package main

import (
	"fmt"
)

const initialBucketSize = 4
const loadFactor = 0.75

// Entry represents a key-value pair.
type Entry struct {
	key   string
	value interface{}
	next  *Entry
}

// HashMap is a resizable hash map.
type HashMap struct {
	buckets []*Entry
	size    int // number of key-value pairs
}

// NewHashMap creates a new HashMap with initial capacity.
func NewHashMap() *HashMap {
	return &HashMap{
		buckets: make([]*Entry, initialBucketSize),
	}
}

// hash function
func hash(key string, bucketCount int) int {
	hash := 0
	for _, ch := range key {
		hash = (hash*31 + int(ch)) % bucketCount
	}
	return hash
}

// Put inserts or updates a key-value pair.
func (m *HashMap) Put(key string, value interface{}) {
	if float64(m.size)/float64(len(m.buckets)) >= loadFactor {
		m.resize()
	}

	index := hash(key, len(m.buckets))
	head := m.buckets[index]

	for current := head; current != nil; current = current.next {
		if current.key == key {
			current.value = value
			return
		}
	}

	newEntry := &Entry{key: key, value: value, next: head}
	m.buckets[index] = newEntry
	m.size++
}

// Get retrieves a value by key.
func (m *HashMap) Get(key string) (interface{}, bool) {
	index := hash(key, len(m.buckets))
	for current := m.buckets[index]; current != nil; current = current.next {
		if current.key == key {
			return current.value, true
		}
	}
	return nil, false
}

// resize doubles the bucket array and rehashes all entries.
func (m *HashMap) resize() {
	newBucketCount := len(m.buckets) * 2
	newBuckets := make([]*Entry, newBucketCount)

	for _, head := range m.buckets {
		for current := head; current != nil; current = current.next {
			newIndex := hash(current.key, newBucketCount)
			newEntry := &Entry{key: current.key, value: current.value, next: newBuckets[newIndex]}
			newBuckets[newIndex] = newEntry
		}
	}
	m.buckets = newBuckets
	// size remains the same
}

// Example usage
func main() {
	hm := NewHashMap()
	for i := 0; i < 20; i++ {
		hm.Put(fmt.Sprintf("key%d", i), i)
	}

	for i := 0; i < 20; i++ {
		if val, ok := hm.Get(fmt.Sprintf("key%d", i)); ok {
			fmt.Printf("key%d: %v\n", i, val)
		}
	}
}
