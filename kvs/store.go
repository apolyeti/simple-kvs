package kvs

import (
	"errors"
)

type Store struct {
	// use maps package to store key-value pairs
	data map[string]int
}

func New() *Store {
	return &Store{data: make(map[string]int)}
}

func (s *Store) Set(key string, value int) {
	s.data[key] = value
}

func (s *Store) Get(key string) (int, error) {
	value, ok := s.data[key]
	if !ok {
		return -1, errors.New("KEY NOT FOUND")
	}

	return value, nil
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}
