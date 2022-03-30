package storage

import (
	"errors"
	"sync"

	"github.com/KokoulinM/go-musthave-shortener-tpl/internal/app/helpers"
)

type Repository interface {
	LinkBy(sl string) (string, error)
	Save(url string) (sl string)
}

type Storage struct {
	data map[string]string
	mu   sync.Mutex
}

func (s *Storage) LinkBy(sl string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	link, ok := s.data[sl]
	if !ok {
		return link, errors.New("url not found")
	}

	return link, nil
}

func (s *Storage) Save(url string) (sl string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	sl = string(helpers.RandomString(10))

	s.data[sl] = url
	return
}

func New() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}
