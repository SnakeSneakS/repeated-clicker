package internal

import (
	"context"
	"sync"
)

type ctxCancelStore struct {
	mu      *sync.Mutex
	counter int
	cancels map[int]context.CancelFunc
}

func newCtxCancelStore() *ctxCancelStore {
	return &ctxCancelStore{
		mu:      new(sync.Mutex),
		counter: 0,
		cancels: make(map[int]context.CancelFunc),
	}
}

func (s *ctxCancelStore) Add(cancel context.CancelFunc) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counter++
	s.cancels[s.counter] = cancel
	return s.counter
}

func (s *ctxCancelStore) Del(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.cancels, id)
}

func (s *ctxCancelStore) Get(id int) (context.CancelFunc, bool) {
	cancel, ok := s.cancels[id]
	return cancel, ok
}

func (s *ctxCancelStore) GetIDs() []int {
	keys := make([]int, len(s.cancels))
	i := 0
	for k := range s.cancels {
		keys[i] = k
		i++
	}
	return keys
}
