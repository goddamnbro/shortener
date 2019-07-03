package shortener

import (
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
)

type Storage map[string]string

//type Shortener interface {
//	Shorten(string) string
//	Resolve(string) string
//}

func (s Storage) Shorten(url string) string {
	h := fnv.New32()
	_, err := h.Write([]byte(url))
	if err != nil {
		log.Println("error:", err)
		return ""
	}

	hash := fmt.Sprintf("%d", h.Sum32())

	for {
		if _, existed := s[hash]; existed {
			hash = fmt.Sprintf("%s%d", hash, rand.Intn(10))
		} else {
			break
		}
	}

	s[hash] = url
	return hash
}

func (s Storage) Resolve(hash string) string {
	return s[hash]
}
