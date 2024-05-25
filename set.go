package dsa

import (
	"errors"
	"fmt"
)

type Set[T comparable] struct {
	items map[T]struct{}
}

// Returns a Set instance
func NewSet[T comparable]() *Set[T] {
	var set Set[T]

	set.items = make(map[T]struct{})
	return &set
}

// Insert - inserts an element to our Set
func (s *Set[T]) Insert(element T) {
	s.items[element] = struct{}{}
}

// Delete - removes an element from our Set if exists
func (s *Set[T]) Delete(element T) error {
	if _, exists := s.items[element]; !exists {
		return errors.New("element not present in Set")
	}
	delete(s.items, element)
	return nil
}

// Contains - returns true if element exist
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.items[element]
	return exists
}

// Size - returns size of Set
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Print set
func (s *Set[T]) List() {
	for k, _ := range s.items {
		fmt.Printf("%v ", k)
	}
	fmt.Println()
}
