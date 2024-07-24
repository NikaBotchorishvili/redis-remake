package internal

import (
	"fmt"
	"sync"
)

type Store struct {
	Data map[string]string
	Mu   sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		Data: make(map[string]string),
	}
}

func (s *Store) Set(key, value string) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Data[key] = value
}

func (s *Store) Get(key string) (string, error){
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	value, exists := s.Data[key]
	if !exists {
		return "", fmt.Errorf("key not found")
	}
	return value, nil
}

func (s *Store) Delete(key string) (string, error){


	return "", nil
}
