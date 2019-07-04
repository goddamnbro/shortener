package shortener

import (
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
)

func (s *Storage) Shorten(url string) (string, error) {
	h := fnv.New32()
	_, err := h.Write([]byte(url))
	if err != nil {
		log.Println("error:", err)
		return "", err
	}

	hash := fmt.Sprintf("%d", h.Sum32())

	for {
		if _, existed := s.Load(hash); existed {
			hash = fmt.Sprintf("%s%d", hash, rand.Intn(10))
		} else {
			break
		}
	}

	s.Store(hash, url)
	return hash, nil
}

func (s *Storage) Resolve(hash string) (string, bool) {
	value, ok := s.Load(hash)
	return value, ok
}
