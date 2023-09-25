package internal

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

type ctxCancelStore struct {
	mu      *sync.Mutex
	cancels map[uuid.UUID]context.CancelFunc
}

func newCtxCancelStore() *ctxCancelStore {
	return &ctxCancelStore{
		mu:      new(sync.Mutex),
		cancels: make(map[uuid.UUID]context.CancelFunc),
	}
}

func (s *ctxCancelStore) Add(cancel context.CancelFunc) uuid.UUID {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New()
	for {
		if _, ok := s.cancels[id]; !ok {
			break
		}
		id = uuid.New()
	}
	s.cancels[id] = cancel
	return id
}

func (s *ctxCancelStore) Del(id uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.cancels, id)
}

func (s *ctxCancelStore) Get(id uuid.UUID) (context.CancelFunc, bool) {
	cancel, ok := s.cancels[id]
	return cancel, ok
}

func (s *ctxCancelStore) GetIDs() []uuid.UUID {
	keys := make([]uuid.UUID, len(s.cancels))
	i := 0
	for k := range s.cancels {
		keys[i] = k
		i++
	}
	return keys
}
