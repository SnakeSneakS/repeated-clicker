package internal

import (
	"context"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
)

func stopRepeatedClick(s *ctxCancelStore) {
	log.Printf("stop repeated clicks\n")

	ids := s.GetIDs()
	for _, id := range ids {
		cancel, ok := s.Get(id)
		if ok {
			cancel()
			s.Del(id)
		}
	}
}

func fireRepeatedClick(s *ctxCancelStore, interval, duration time.Duration) {
	log.Printf("repeatedly clicks for %s with interval %s\n", duration.String(), interval.String())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cancelID := s.Add(cancel)
	defer s.Del(cancelID)

	endTime := time.Now().Add(duration)
outer:
	for time.Now().Before(endTime) {
		select {
		case <-ctx.Done():
			break outer
		default:
			robotgo.Click()
			time.Sleep(interval)
			continue outer
		}
	}

}
