package shortener

import (
	"sync"
)

// I'm sure, we don't have cache contention issue :)
type Storage struct {
	mx sync.Mutex
	m  map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		m: make(map[string]string),
	}
}

func (s *Storage) Load(key string) (string, bool) {
	s.mx.Lock()
	defer s.mx.Unlock()
	val, ok := s.m[key]
	return val, ok
}

func (s *Storage) Store(key string, value string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.m[key] = value
}

func (s *Storage) Range() map[string]string {
	s.mx.Lock()
	defer s.mx.Unlock()

	newMap := make(map[string]string, len(s.m))

	for k, v := range s.m {
		newMap[k] = v
	}

	return newMap
}
