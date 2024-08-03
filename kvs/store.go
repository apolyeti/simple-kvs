package kvs

import (
	"errors"
	"sync"
)

type Store struct {
	// use maps package to store key-value pairs
	data map[string]int
	// use sync
	mu sync.RWMutex
}

func New() *Store {
	return &Store{data: make(map[string]int)}
}

func (s *Store) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *Store) Get(key string) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	if !ok {
		return -1, errors.New("KEY NOT FOUND")
	}

	return value, nil
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}
